package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	// 移除数据库支持
	//InitModel()
}

func InitModel() {
	DB, err = gorm.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 设置
	DB.SingularTable(true)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	//
	//if config.Config.ENV == "dev" {
	//	DB = DB.Debug()
	//}
	DB = DB.Debug()
	log.Println("Init DB Complete")
}

func InitTables() {
	//DB.AutoMigrate(dto.AnkiNoteInfo{})
}
