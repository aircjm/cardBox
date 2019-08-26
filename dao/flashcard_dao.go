package dao

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/model/request"
	"log"
)

// 新增或者更新trello_card数据
func SaveCardOrm(card trello.Card) {
	oldFlashCard := dto.FlashCard{}
	oldFlashCard.ID = card.ID
	old := DB.First(&oldFlashCard)
	if old != nil {
		log.Println("更新FlashCard")
		oldFlashCard.SetFlashCard(card)
		DB.Model(&oldFlashCard).Updates(&oldFlashCard)
	} else {
		log.Println("新增FlashCard")
		oldFlashCard.NewFlashCard(card)
		DB.Create(&oldFlashCard)
	}
}

//GetCardByCardIdList 通过卡片id集合获取卡片
func GetCardByCardIdList(cardIdList []string) []dto.FlashCard {
	var flashCardList []dto.FlashCard
	DB.Where("id in (?)", cardIdList).Find(&flashCardList)
	return flashCardList
}

// 获取更新dto.FlashCard 数据 通过主键id获取
func GetCardList(request request.GetCardListRequest) []dto.FlashCard {
	cards := []dto.FlashCard{}
	where := ""
	if request.HaveAnki > 0 {
		where = where + ""
	}

	db := DB.Where(where)
	db.Find(&cards)
	return cards
}
