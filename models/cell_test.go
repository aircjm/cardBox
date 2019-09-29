package models

import (
	"github.com/aircjm/gocard/client"
	"testing"
)

func TestCreateOrUpdateCell(t *testing.T) {

	card := client.GetTestCard()
	cell := new(Cell).ConvertToCell(card)
	t.Log(cell)
	CreateOrUpdateCell(&cell)
}
