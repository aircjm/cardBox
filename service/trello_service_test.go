package service

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/client"
	"log"
	"testing"
)

func TestSaveAllCards(t *testing.T) {
	SaveAllCards()
}

func TestSaveAllCardsOrm(t *testing.T) {
	boards, err := client.TrelloCL.GetMyBoards(trello.Defaults())
	if err != nil {
		log.Fatal(err)
	}

	for _, board := range boards {
		cards, err := board.GetCards(trello.Defaults())
		if err != nil {
			log.Fatal(err)
		}
		go SaveCardsOrm(cards)
	}
}
