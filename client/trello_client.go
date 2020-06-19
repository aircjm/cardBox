package client

import (
	"github.com/adlio/trello"
	"log"
)

// 获取trello的client
var TrelloCL = trello.NewClient("", "")

func GetTestCard() *trello.Card {
	boards, err := TrelloCL.GetMyBoards(trello.Defaults())
	cards, err := boards[0].GetCards(trello.Defaults())
	if err != nil {
		log.Fatalln(err)
	}
	return cards[0]
}
