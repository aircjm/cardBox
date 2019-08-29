package dao

import (
	"github.com/aircjm/gocard/dto"
	"log"
)

func SaveAnkiNote(ankiNote dto.AnkiNoteInfo) {
	info := dto.AnkiNoteInfo{}
	DB.Find("anki_note_id is ?", ankiNote.AnkiNoteID).First(&info)
	if info.ID > 0 {
		// 更新info
		DB.Update(&info)
	} else {
		DB.Save(&ankiNote)
		log.Println(ankiNote.ID)
	}

}
