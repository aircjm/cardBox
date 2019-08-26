package service

import (
	"github.com/aircjm/gocard/client"
	"github.com/aircjm/gocard/client/model"
	"testing"
)

func TestAddAnkiNote(t *testing.T) {

	card := client.GetTestCard()
	addNoteAnkiRequest := model.AnkiAddNoteRequest{}.GetAnkiAddNote(card)

	noteId := AddAnkiNote(*addNoteAnkiRequest)

	t.Log(noteId)
}
