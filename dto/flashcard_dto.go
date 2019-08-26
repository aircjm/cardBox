package dto

import (
	"github.com/adlio/trello"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type FlashCard struct {
	ID           string `gorm:"primary_key"`
	Name         string `gorm:"not null";index`
	Desc         string
	AnkiStatus   int // 0 表示待处理 -1 标识放弃不生成anki笔记 1标识生成anki
	CardType     int
	TrelloCardB  postgres.Jsonb `gorm:"type:jsonb;"`
	TrelloCard   trello.Card    `gorm:"-"`
	AnkiNoteInfo AnkiNoteInfo   `gorm:ForeignKey:ID;AssociationForeignKey:Refer`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Closed       int
}

func (FlashCard) NewFlashCard(trelloCard trello.Card) *FlashCard {
	flashCard := FlashCard{}
	flashCard.ID = trelloCard.ID
	flashCard.CardType = 1
	flashCard.Name = trelloCard.Name
	flashCard.Desc = trelloCard.Desc
	flashCard.AnkiStatus = 0
	flashCard.CreatedAt = time.Now()
	flashCard.UpdatedAt = time.Now()
	return &flashCard
}

func (FlashCard) SetFlashCard(trelloCard trello.Card) *FlashCard {
	flashCard := FlashCard{}
	flashCard.ID = trelloCard.ID
	flashCard.Name = trelloCard.Name
	flashCard.Desc = trelloCard.Desc
	flashCard.UpdatedAt = time.Now()
	return &flashCard
}
