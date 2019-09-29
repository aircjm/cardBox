package dto

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" gorm: unique_index`
	Password string `json:"password"`
}
