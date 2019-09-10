package cardDao

import (
	"github.com/adlio/trello"
	"github.com/aircjm/gocard/dao"
	"github.com/aircjm/gocard/dto"
	"github.com/aircjm/gocard/model/request"
	"log"
)

// 新增或者更新trello_card数据
func SaveCardOrm(card trello.Card) {
	oldFlashCard := dto.FlashCard{}
	oldFlashCard.ID = card.ID
	count := 0
	dao.DB.Model(&dto.FlashCard{}).Where("id = ?", card.ID).Count(&count)
	if count > 0 {
		log.Println("更新FlashCard")
		flashCard := dto.FlashCard{}.UpdateFlashCard(card)
		dao.DB.Model(&dto.FlashCard{}).Updates(&flashCard)
	} else {
		log.Println("新增FlashCard")
		flashCard := dto.FlashCard{}.NewFlashCard(card)
		dao.DB.Create(&flashCard)
	}
}

func UpdateCard(card dto.FlashCard) {
	dao.DB.Model(&card).Update("card_status", card.CardStatus)
}

//GetCardao.DByCardIdList 通过卡片id集合获取卡片
func GetCardByCardIdList(cardIdList []string) []dto.FlashCard {
	var flashCardList []dto.FlashCard
	dao.DB.Where("id in (?)", cardIdList).Find(&flashCardList)
	for index := range flashCardList {
		var ankiNote dto.AnkiNoteInfo
		dao.DB.Where("trello_card_id = ?", flashCardList[index].ID).First(&ankiNote)
		flashCardList[index].AnkiNoteInfo = ankiNote
	}
	return flashCardList
}

func SaveBoard(board trello.Board) {

	oldMingBoard := dto.MingBoard{}
	oldMingBoard.ID = board.ID
	dao.DB.Where("id = ?", oldMingBoard.ID).First(&oldMingBoard)
	if oldMingBoard.Name != "" {
		log.Println("更新board")
		mingBoard := oldMingBoard.SetMingBoard(board)
		dao.DB.Model(&mingBoard).Updates(&mingBoard)
	} else {
		log.Println("新增board")
		mingBoard := oldMingBoard.NewMingBoard(board)
		dao.DB.Create(&mingBoard)
	}
}

//GetBoardList 获取所有的boardList
func GetBoardList() []dto.MingBoard {
	var boards []dto.MingBoard
	dao.DB.Find(&boards)
	return boards
}

//GetBoardList 获取所有的boardList
func GetCardList(request request.GetCardListRequest) ([]dto.FlashCard, int) {
	var cards []dto.FlashCard
	var count = 0
	db := dao.DB
	if request.CardStatus > 0 {
		db = db.Where("card_status = ?", request.CardStatus)
	}
	if len(request.BoardId) > 0 {
		db = db.Where("id_board = ?", request.BoardId)
	}

	db.Order("id desc")
	db.Model(&dto.FlashCard{}).Count(&count)
	if request.Pagination.PageSize >= 0 {
		//Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
		db = db.Limit(request.Pagination.PageSize).Offset((request.Pagination.CurrentPage - 1) * request.Pagination.PageSize)
	}
	db.Find(&cards)
	log.Println(count)
	return cards, count
}

func GetBoardListByBoardIdList(boardIdList []string) []dto.MingBoard {
	var boardList []dto.MingBoard
	dao.DB.Where("id in (?)", boardIdList).Find(&boardList)
	return boardList
}
