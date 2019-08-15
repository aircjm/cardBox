package client

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/config"
	"log"
)

// 获取trello的client
var TrelloCL = trello.NewClient(config.TrelloAPI, config.TrelloToken)

func GetTestCard() *trello.Card {
	boards, err := TrelloCL.GetMyBoards(trello.Defaults())
	cards, err := boards[0].GetCards(trello.Defaults())
	if err != nil {
		log.Fatalln(err)
	}
	return cards[0]
}
