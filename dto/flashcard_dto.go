package dto

import (
	"github.com/adlio/trello"
	"github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type FlashCard struct {
	ID               string         `json:"id" gorm:"primary_key"`
	Name             string         `json:"name" gorm:"not null";index`
	Desc             string         `json:"desc"`
	CardType         int            `json:"cardType"`
	TrelloCardB      postgres.Jsonb `"gorm:"type:jsonb;"`
	TrelloCard       trello.Card    `gorm:"-"`
	AnkiNoteInfo     AnkiNoteInfo   `json:"ankiNoteInfo" gorm:ForeignKey:ID;AssociationForeignKey:Refer`
	DateLastActivity time.Time      `json:"dateLastActivity"`
	CardStatus       int            `json:"cardStatus"` // 0 表示待处理 -1 标识放弃不生成anki笔记 1标识生成anki
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	Closed           int            `json:"closed"`
	IDBoard          string         `json:"idBoard"`
}

type MingBoard struct {
	ID             string    `json:"id" gorm:"primary_key"`
	Name           string    `json:"name" gorm:"not null";index`
	Desc           string    `json:"desc"`
	Closed         bool      `json:"closed"`
	IDOrganization string    `json:"idOrganization"`
	Pinned         bool      `json:"pinned"`
	URL            string    `json:"url"`
	ShortURL       string    `json:"shortUrl"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (MingBoard) NewMingBoard(board trello.Board) *MingBoard {
	mingBoard := MingBoard{}
	mingBoard.ID = board.ID
	mingBoard.Name = board.Name
	mingBoard.Desc = board.Desc
	mingBoard.Closed = board.Closed
	mingBoard.IDOrganization = board.IDOrganization
	mingBoard.Pinned = board.Pinned
	mingBoard.URL = board.URL
	mingBoard.ShortURL = board.ShortURL
	return &mingBoard
}

func (MingBoard) SetMingBoard(board trello.Board) *MingBoard {
	mingBoard := MingBoard{}
	mingBoard.ID = board.ID
	mingBoard.Name = board.Name
	mingBoard.Desc = board.Desc
	mingBoard.Closed = board.Closed
	mingBoard.IDOrganization = board.IDOrganization
	mingBoard.Pinned = board.Pinned
	mingBoard.URL = board.URL
	mingBoard.ShortURL = board.ShortURL
	mingBoard.UpdatedAt = time.Now()
	return &mingBoard
}

func (FlashCard) NewFlashCard(trelloCard trello.Card) *FlashCard {
	flashCard := FlashCard{}
	flashCard.ID = trelloCard.ID
	flashCard.CardType = 1
	flashCard.Name = trelloCard.Name
	flashCard.Desc = trelloCard.Desc
	flashCard.CardStatus = 0
	flashCard.CreatedAt = time.Now()
	flashCard.UpdatedAt = time.Now()
	flashCard.DateLastActivity = *trelloCard.DateLastActivity
	return &flashCard
}

func (FlashCard) UpdateFlashCard(trelloCard trello.Card) *FlashCard {
	flashCard := FlashCard{}
	flashCard.ID = trelloCard.ID
	flashCard.Name = trelloCard.Name
	flashCard.Desc = trelloCard.Desc
	flashCard.IDBoard = trelloCard.IDBoard
	flashCard.UpdatedAt = time.Now()
	flashCard.DateLastActivity = *trelloCard.DateLastActivity
	return &flashCard
}
