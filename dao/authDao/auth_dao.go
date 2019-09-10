package authDao

import (
	"github.com/aircjm/gocard/dao"
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/util"
)

func CheckAuth(username, password string) dto.User {
	var auth dto.User
	dao.DB.Select("id").Where(dto.User{Username: username, Password: util.EncodeMD5(password)}).First(&auth)
	if auth.ID > 0 {
		return auth
	}
	return dto.User{}
}

func SaveUser(user dto.User) {
	// insert or update
	count := 0
	dao.DB.Model(&user).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		dao.DB.Model(&user).Where("username = ?", user.Username).Update("password", user.Password).Limit(1)
	} else {
		dao.DB.Model(&user).Create(&user)
	}
}
