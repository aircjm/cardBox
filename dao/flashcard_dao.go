package dao

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/model/request"
	"log"
)

// 新增或者更新trello_card数据
func SaveCard(card trello.Card) {
	oldCard := GetCardByCardId(card.ID)
	log.Println("当前Id: ", card.ID, "数据库id为：", oldCard.TrelloCard.ID)
	if len(oldCard.TrelloCard.ID) != 0 {
		oldCard.TrelloCard = card
		log.Println("开始更新")
		DB.Exec("update json_card set attrs = $2 where id = $1", card.ID, oldCard)
	} else {
		log.Println("开始新增")
		newCard := dto.FlashCard{}.SetFlashCard(card)
		DB.Exec("INSERT INTO json_card (id,attrs) VALUES ($1, $2)", card.ID, newCard)
	}
}

// 新增或者更新trello_card数据
func SaveCardOrm(card trello.Card) {
	oldFlashCard := dto.FlashCard{}
	oldFlashCard.ID = card.ID
	old := DB.First(&oldFlashCard)
	if old != nil {
		log.Println("更新FlashCard")
		oldFlashCard.SetFlashCard(card)
		DB.Save(oldFlashCard)
	} else {
		log.Println("新增FlashCard")
		oldFlashCard.NewFlashCard(card)
		DB.Create(oldFlashCard)
	}

}

// 获取更新dto.FlashCard 数据 通过主键id获取
func GetCardByCardId(cardId string) dto.FlashCard {
	var result dto.FlashCard
	raw := DB.Raw("SELECT attrs FROM json_card WHERE id =  $1", cardId)
	raw.Row().Scan(&result)
	return result
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
