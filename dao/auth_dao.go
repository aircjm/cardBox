package dao

import (
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/util"
)

func CheckAuth(username, password string) bool {
	var auth dto.User
	DB.Select("id").Where(dto.User{Username: username, Password: util.EncodeMD5(password)}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func InsertUser(user dto.User) {

}
