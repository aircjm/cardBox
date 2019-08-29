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

func TestSaveCardsOrm(t *testing.T) {
	card := client.GetTestCard()
	var cards []*trello.Card
	cards = append(cards, card)
	SaveCardsOrm(cards)
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

func TestSaveRecentlyEditedCard(t *testing.T) {
	SaveRecentlyEditedCard()
}

func TestSaveBoard(t *testing.T) {
	boards, err := client.TrelloCL.GetMyBoards(trello.Defaults())
	if err != nil {

	}

	for _, board := range boards {
		SaveBoard(board)
	}

}
