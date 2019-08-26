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
	IDBoard      string `json:"idBoard"`
}

type MingBoard struct {
	ID             string `gorm:"primary_key"`
	Name           string `gorm:"not null";index`
	Desc           string
	Closed         bool
	IDOrganization string `json:"idOrganization"`
	Pinned         bool   `json:"pinned"`
	URL            string `json:"url"`
	ShortURL       string `json:"shortUrl"`
	UpdatedAt      time.Time
}

func (MingBoard) NewMingBoard(board trello.Board) *MingBoard {
	mingBoard := MingBoard{}
	mingBoard.ID = board.ID
	mingBoard.Name = board.Name
	mingBoard.Desc = board.Desc
	mingBoard.Closed = board.Closed
	mingBoard.IDOrganization = board.IDOrganization
	mingBoard.IDOrganization = board.IDOrganization
	mingBoard.Pinned = board.Pinned
	mingBoard.URL = board.URL
	mingBoard.ShortURL = board.ShortURL
	return &mingBoard
}

func (MingBoard) SetMingBoardd(board trello.Board) *MingBoard {
	mingBoard := MingBoard{}
	mingBoard.ID = board.ID
	mingBoard.Name = board.Name
	mingBoard.Desc = board.Desc
	mingBoard.UpdatedAt = time.Now()
	return &mingBoard
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
	flashCard.IDBoard = trelloCard.IDBoard
	flashCard.UpdatedAt = time.Now()
	return &flashCard
}
