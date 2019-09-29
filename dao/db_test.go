package dao

import (
	"github.com/aircjm/gocard/dto"
	"testing"
)

func TestGetDB(t *testing.T) {
	DB.AutoMigrate(&dto.FlashCard{}, &dto.AnkiNoteInfo{}, &dto.MingBoard{}, &dto.Cell{})
}
