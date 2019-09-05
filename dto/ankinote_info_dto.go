package dto

import "time"

// anki笔记实体
type AnkiNoteInfo struct {
	ID           uint  `gorm:"primary_key"`
	AnkiNoteID   int64 `gorm: unique_index`
	DeckName     string
	ModelName    string
	Title        string
	HtmlContext  string
	TrelloCardId string `gorm: unique_index`
	Status       uint   `gorm:"default:0"` // 0 表示待处理 1 表示已生成
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
