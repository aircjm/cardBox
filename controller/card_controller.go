package controller

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/model/request"
	"github.com/aircjm/gocard/service"
	"github.com/gin-gonic/gin"
	"log"
)

type BatchCardOpt struct {
	TrelloCardIds []string // cardId集合
	HandlerType   int      // 操作类型
	HandlerName   string   // 操作名称
}

func GetCardByCardId(c *gin.Context) {
	//cardId := c.Param("cardId")
	cG := common.Gin{C: c}
	//mingCard := service.GetMingCard(cardId)
	cG.Response(200, 0, nil)

}

// 设置card到anki卡片中
func SaveCardToAnki(c *gin.Context) {
	cG := common.Gin{C: c}
	opt := BatchCardOpt{}
	err := c.BindJSON(&opt)
	if err != nil {
		log.Fatalln(err)
	}
	service.SaveCardToAnki(opt.TrelloCardIds)
	cG.Response(200, 0, nil)

}

func SaveRecentCard(c *gin.Context) {
	cG := common.Gin{C: c}
	service.SaveRecentlyEditedCard()
	cG.Response(200, 0, nil)

}

func GetRecentCard(c *gin.Context) {
	cG := common.Gin{C: c}
	cards, err := service.GetRecentlyEditedCard()
	if err != nil {
		log.Fatalln(err)
	}
	cG.Response(200, 0, cards)
}

func GetCardList(c *gin.Context) {
	cG := common.Gin{C: c}

	request := request.GetCardListRequest{}
	c.BindJSON(&request)
	cards, err := service.GetCardList(request)
	if err != nil {
		log.Fatalln(err)
	}
	cG.Response(200, 0, cards)
}

func SaveAllCards(c *gin.Context) {
	cG := common.Gin{C: c}
	service.SaveAllCards()
	cG.Response(200, 0, nil)
}
