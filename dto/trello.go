package dto

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type TrelloEntity struct {
	gorm.Model
	ID   string `gorm:"primary_key"`
	Name string
	Type int
}

func (c TrelloEntity) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *TrelloEntity) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
