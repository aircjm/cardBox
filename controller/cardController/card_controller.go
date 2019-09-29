package cardController

import (
	"github.com/aircjm/gocard/common"
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/model/request"
	"github.com/aircjm/gocard/model/response"
	"github.com/aircjm/gocard/service/ankiService"
	"github.com/aircjm/gocard/service/trelloService"
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
	ankiService.SaveCardToAnki(opt.TrelloCardIds)
	cG.Response(200, 0, nil)

}

func GetRecentCard(c *gin.Context) {
	cG := common.Gin{C: c}
	cards, err := trelloService.GetRecentlyEditedCard()
	if err != nil {
		log.Fatalln(err)
	}
	cG.Response(200, 0, cards)
}

func GetCardList(c *gin.Context) {
	cG := common.Gin{C: c}

	request := request.GetCardListRequest{}
	c.BindJSON(&request)
	cards, count := ankiService.GetCardList(request)
	cG.Response(200, 0, response.ListResponse{Count: count, List: cards})
}

func SaveAllCards(c *gin.Context) {
	cG := common.Gin{C: c}
	trelloService.SaveAllCards()
	cG.Response(200, 0, nil)
}

//ConvertToAnki 将卡片转换成anki的note
func ConvertToAnki(c *gin.Context) {
	cG := common.Gin{C: c}
	request := request.CardIdList{}
	c.BindJSON(&request)
	trelloService.ConvertToAnki(request.CardIdList)
	cG.Response(200, 0, nil)
}

func UpdateCardStatus(c *gin.Context) {
	cG := common.Gin{C: c}
	request := dto.FlashCard{}
	c.BindJSON(&request)

	trelloService.UpdateCardStatus(request)

	cG.Response(200, 0, nil)
}

func SaveRecentCard(c *gin.Context) {
	cG := common.Gin{C: c}
	trelloService.SaveRecentlyEditedCard()
	cG.Response(200, 0, nil)

}
