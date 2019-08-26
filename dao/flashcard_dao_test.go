package dao

import (
	"github.com/aircjm/gocard/client"
	"testing"
)

func TestSaveCardOrm(t *testing.T) {
	card := client.GetTestCard()
	SaveCardOrm(*card)
}

// 单元测试
func TestGetCardByCardIdList(t *testing.T) {
	cardId := client.GetTestCard().ID
	var cardIdList []string
	cardIdList = append(cardIdList, cardId)
	flashCards := GetCardByCardIdList(cardIdList)

	if len(flashCards) > 0 {
		t.Log(flashCards[0])
	} else {
		t.Fatal("查不到数据")
	}
}
