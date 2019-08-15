package dao

import (
	"github.com/aircjm/gocard/client"
	"testing"
)

func TestSaveCardOrm(t *testing.T) {
	card := client.GetTestCard()
	SaveCardOrm(*card)
}
