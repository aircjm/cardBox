package dao

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/dto"
	"log"
)

// 新增或者更新trello_card数据
func SaveCardOrm(card trello.Card) {
	oldFlashCard := dto.FlashCard{}
	oldFlashCard.ID = card.ID
	old := DB.First(&oldFlashCard)
	if old != nil {
		log.Println("更新FlashCard")
		flashCard := oldFlashCard.SetFlashCard(card)
		DB.Model(&flashCard).Updates(&flashCard)
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
	for index := range flashCardList {
		var ankiNote dto.AnkiNoteInfo
		DB.Where("trello_card_id = ?", flashCardList[index].ID).First(&ankiNote)
		flashCardList[index].AnkiNoteInfo = ankiNote
	}
	return flashCardList
}

func SaveBoard(board trello.Board) {

	oldMingBoard := dto.MingBoard{}
	oldMingBoard.ID = board.ID
	DB.Where("id = ?", oldMingBoard.ID).First(&oldMingBoard)
	if oldMingBoard.Name != "" {
		log.Println("更新board")
		mingBoard := oldMingBoard.SetMingBoard(board)
		DB.Model(&mingBoard).Updates(&mingBoard)
	} else {
		log.Println("新增board")
		mingBoard := oldMingBoard.NewMingBoard(board)
		DB.Create(&mingBoard)
	}
}

//GetBoardList 获取所有的boardList
func GetBoardList() []dto.MingBoard {
	var boards []dto.MingBoard
	DB.Find(&boards)
	return boards
}
