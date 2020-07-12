package dao

import (
	"github.com/aircjm/cardBox/dto"
	"testing"
)

func TestGetDB(t *testing.T) {
	DB.AutoMigrate(&dto.FlashCard{}, &dto.AnkiNoteInfo{}, &dto.TrelloEntity{})
}
