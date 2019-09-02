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
	count := 0
	DB.Model(&dto.FlashCard{}).Where("id = ?", card.ID).Count(&count)
	if count > 0 {
		log.Println("更新FlashCard")
		flashCard := dto.FlashCard{}.SetFlashCard(card)
		DB.Model(&dto.FlashCard{}).Updates(&flashCard)
	} else {
		log.Println("新增FlashCard")
		flashCard := dto.FlashCard{}.NewFlashCard(card)
		DB.Create(&flashCard)
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

//GetBoardList 获取所有的boardList
func GetCardList(request request.GetCardListRequest) ([]dto.FlashCard, int) {
	var cards []dto.FlashCard
	var count = 0
	db := DB
	if request.CardStatus >= 0 {
		db = db.Where("card_status = ?", request.CardStatus)
	}
	if len(request.BoardId) > 0 {
		db = db.Where("id_board = ?", request.BoardId)
	}

	if request.Pagination.PageSize >= 0 {
		//Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
		db = db.Limit(request.Pagination.PageSize).Offset((request.Pagination.CurrentPage - 1) * request.Pagination.PageSize)
	}
	db.Find(&cards).Count(&count)
	log.Println(count)
	return cards, count
}

func GetBoardListByBoardIdList(boardIdList []string) []dto.MingBoard {
	var boardList []dto.MingBoard
	DB.Where("id in (?)", boardIdList).Find(&boardList)
	return boardList
}
