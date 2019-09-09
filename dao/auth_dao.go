package dao

import "github.com/aircjm/gocard/dto"

func CheckAuth(username, password string) bool {
	var auth dto.Auth
	DB.Select("id").Where(dto.Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
