package dto

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/adlio/trello"
)

type TrelloEntity struct {
	ID    string
	Name  string
	Type  int
	Board trello.Board `sql:"TYPE:json"`
	Card  trello.Card  `sql:"TYPE:json"`
}

func (c TrelloEntity) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TrelloEntity) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
