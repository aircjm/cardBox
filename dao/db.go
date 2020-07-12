package dao

import (
	"fmt"
	"github.com/aircjm/cardBox/config"
	"github.com/aircjm/cardBox/dto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"os"
	"strings"
)

var DB *gorm.DB
var err error

func Init() (*gorm.DB, error) {
	Conf := config.Conf
	db, err := openDB(*Conf)
	if err != nil {
		return nil, errors.Wrap(err, "open database")
	}

	return db, db.DB().Ping()
}

func init() {
	DB, err = Init()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 设置
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	if config.DEV == "dev" {
		DB = DB.Debug()
	}

	DB.AutoMigrate(&dto.FlashCard{}, &dto.AnkiNoteInfo{}, &dto.TrelloEntity{})

	log.Println("Init DB Complete")
}

func openDB(opts config.Config) (*gorm.DB, error) {
	dsn, err := parseDSN(opts)
	if err != nil {
		return nil, errors.Wrap(err, "parse DSN")
	}

	return gorm.Open(opts.DataBase.Type, dsn)
}

func parseDSN(opts config.Config) (dsn string, err error) {
	// In case the database name contains "?" with some parameters

	databaseConfig := opts.DataBase
	concate := "?"
	if strings.Contains(databaseConfig.Name, concate) {
		concate = "&"
	}

	switch databaseConfig.Type {
	case "mysql":
		if databaseConfig.Host[0] == '/' { // Looks like a unix socket
			dsn = fmt.Sprintf("%s:%s@unix(%s)/%s%scharset=utf8mb4&parseTime=true",
				databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Name, concate)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8mb4&parseTime=true",
				databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Name, concate)
		}

	case "postgres":
		host, port := parsePostgreSQLHostPort(databaseConfig.Host)
		if host[0] == '/' { // looks like a unix socket
			dsn = fmt.Sprintf("postgres://%s:%s@:%s/%s%ssslmode=%s&host=%s",
				url.QueryEscape(databaseConfig.User), url.QueryEscape(databaseConfig.Password), port, databaseConfig.Name, concate, databaseConfig.SSLMode, host)
		} else {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
				url.QueryEscape(databaseConfig.User), url.QueryEscape(databaseConfig.Password), host, port, databaseConfig.Name, concate, databaseConfig.SSLMode)
		}

	//case "mssql":
	//	host, port := parseMSSQLHostPort(opts.Host)
	//	dsn = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;",
	//		host, port, opts.Name, opts.User, opts.Password)
	//
	//case "sqlite3":
	//	dsn = "file:" + opts.Path + "?cache=shared&mode=rwc"

	default:
		return "", errors.Errorf("unrecognized dialect: %s", databaseConfig.Type)
	}

	return dsn, nil
}

func parsePostgreSQLHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "5432"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}
	return host, port
}
