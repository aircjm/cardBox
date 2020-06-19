package dto

import (
	"github.com/jinzhu/gorm"
)

// anki笔记实体
type AnkiNoteInfo struct {
	gorm.Model
	TrelloCardId string
	Title        string
	HtmlContext  string
	ModelName    string
}
