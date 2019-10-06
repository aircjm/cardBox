package trelloService

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/client"
	"github.com/aircjm/gocard/dto"
	"log"
	"sync"
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

func TestBatchSaveTrelloCardToCell(t *testing.T) {
	boards, err := client.TrelloCL.GetMyBoards(trello.Defaults())
	if err != nil {

	}
	// 试试看看有没有更好的多线程的方式
	var wg sync.WaitGroup
	for _, board := range boards {
		trelloCards, err := board.GetCards(trello.Defaults())
		if err != nil {
			log.Fatal(err)
		}
		wg.Add(1)
		go wgDone(&wg, trelloCards)

	}
	wg.Wait()
}

func wgDone(wg *sync.WaitGroup, cards []*trello.Card) {
	BatchSaveTrelloCardToCell(cards)
	wg.Done()
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

func TestConvertToAnkiNote(t *testing.T) {
	card := client.GetTestCard()
	var cardIdList []string
	ConvertToAnkiNote(append(cardIdList, card.ID))
}

func TestSingleConvertToAnki(t *testing.T) {
	boards, _ := client.TrelloCL.GetMyBoards(trello.Defaults())
	for _, board := range boards {
		cards, _ := board.GetCards(trello.Defaults())
		for _, card := range cards {
			SingleConvertToAnki(card.ID)
		}

	}
}

func TestUpdateCardStatus(t *testing.T) {

	card := dto.FlashCard{}

	card.ID = "12342131"
	card.CardStatus = 1

	UpdateCardStatus(card)

}
