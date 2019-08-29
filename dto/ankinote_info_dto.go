package dto

import "github.com/jinzhu/gorm"

// anki笔记实体
type AnkiNoteInfo struct {
	gorm.Model
	TrelloCardId string `gorm: unique_index`
	Title        string
	HtmlContext  string
	ModelName    string
	AnkiNoteID   int64 `gorm: unique_index`
	Status       uint  `gorm:"default:0"` // 0 表示待处理 1 表示已生成
}
