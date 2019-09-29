package authDao

import (
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/util"
	"testing"
)

func TestSaveUser(t *testing.T) {
	user := dto.User{Username: "admin", Password: util.EncodeMD5("admin")}
	SaveUser(user)
}
