package service

import (
	"encoding/json"
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/client"
	"github.com/aircjm/gocard/client/model"
	"github.com/aircjm/gocard/config"
	"github.com/aircjm/gocard/dao"
	"github.com/aircjm/gocard/model/request"
	"github.com/aircjm/gocard/model/response"
	"github.com/aircjm/gocard/util"
	"log"
)

type AnkiService interface {
}

func AddAnkiNoteByCardId(cardId string) {
	card, err := client.TrelloCL.GetCard(cardId, trello.Defaults())
	if err != nil {
		panic(err)
	}
	addNoteAnkiRequest := model.AnkiAddNoteRequest{}.GetAnkiAddNote(card)
	response := util.Post(config.AnkiConnect, addNoteAnkiRequest, util.ApplicationJSON)
	log.Println("anki返回的数据是", response)
	ankiResponse := model.AnkiResponse{}
	_ = json.Unmarshal([]byte(response), &ankiResponse)
	log.Println(ankiResponse.Result)
	// 更新anki时间
}

func AddAnkiNote(request model.AnkiAddNoteRequest) int64 {
	response := util.Post(config.AnkiConnect, request, util.ApplicationJSON)
	ankiResponse := model.AnkiResponse{}
	_ = json.Unmarshal([]byte(response), &ankiResponse)
	log.Println(ankiResponse.Result)
	if len(ankiResponse.Error) > 0 {
		log.Println("返回错误，返回信息为：", ankiResponse.Error)
		return 0
	} else {
		return ankiResponse.Result
	}

}

func UpdateAnkiNote(cardId string) {
	card, err := client.TrelloCL.GetCard(cardId, trello.Defaults())
	if err != nil {
		panic(err)
	}
	addNoteAnkiRequest := model.AnkiAddNoteRequest{}.GetAnkiAddNote(card)
	response := util.Post(config.AnkiConnect, addNoteAnkiRequest, util.ApplicationJSON)
	log.Println("anki返回的数据是", response)
	ankiResponse := model.AnkiResponse{}
	_ = json.Unmarshal([]byte(response), &ankiResponse)
	log.Println(ankiResponse.Result)
	// 更新anki时间
}

func TestAnkiConnect() bool {

	response := util.Get(config.AnkiConnect)

	if len(response) != 0 {
		return true
	}

	return false
}

func SaveTrelloToAnkiDecks() {
	boards, err := client.TrelloCL.GetMyBoards(trello.Defaults())

	if err != nil {
		panic(err)
	}

	if !TestAnkiConnect() {
		panic("无法连接anki服务")
	}

	// 需要添加如果已经有deck不要二次添加的逻辑
	for _, board := range boards {
		addDeckRequest := model.AnkiAddDeckRequest{}
		addDeckRequest.Action = "createDeck"
		addDeckRequest.Version = 6
		addDeckRequest.Params.Deck = board.Name
		util.Post(config.AnkiConnect, addDeckRequest, util.ApplicationJSON)
	}
}

// SaveCardToAnki 保存数据到anki
func SaveCardToAnki(Ids []string) {
	for e := range Ids {
		AddAnkiNoteByCardId(Ids[e])
	}
}

func GetCardList(request request.GetCardListRequest) ([]response.CardResponse, int) {
	cardResponseList := []response.CardResponse{}
	cardList, count := dao.GetCardList(request)
	for _, card := range cardList {
		cardResponse := response.CardResponse{}

		cardResponse.CardInfo.Id = card.ID
		cardResponse.CardInfo.Name = card.Name

		cardResponseList = append(cardResponseList, cardResponse)
	}
	return cardResponseList, count
}
