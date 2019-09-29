package ankiDao

import (
	"github.com/aircjm/gocard/dao"
	"github.com/aircjm/gocard/dto"
	"log"
)

func SaveAnkiNote(ankiNote dto.AnkiNoteInfo) {
	info := dto.AnkiNoteInfo{}
	dao.DB.Where("anki_note_id = ?", ankiNote.AnkiNoteID).First(&info)
	if info.ID > 0 {
		log.Println("更新ankinote卡片")
		// 更新info
		dao.DB.Update(&info)
	} else {
		log.Println("新增ankinote卡片")
		dao.DB.Save(&ankiNote)
		if ankiNote.ID > 0 {
			log.Println("新增ankinote卡片成功，id为：", ankiNote.ID)
		}

	}
}

func GetAnkiNoteByTrelloCardId(trelloCardId string) dto.AnkiNoteInfo {
	info := dto.AnkiNoteInfo{}
	dao.DB.Where("trello_card_id = ?", trelloCardId).First(&info)
	return info
}
