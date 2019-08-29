package dto

import "github.com/jinzhu/gorm"

// anki笔记实体
type AnkiNoteInfo struct {
	gorm.Model
	TrelloCardId string `gorm: unique_index`
	Title        string
	HtmlContext  string
	ModelName    string
	AnkiNoteID   uint `gorm: unique_index`
	Status       uint `gorm:"default:0"` // 0 表示待处理
}
