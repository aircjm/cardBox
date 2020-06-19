package response

// Card 基础数据
type CardBaseResponse struct {
	Id   string
	Name string
}

// GetCardList列表返回数据
type CardResponse struct {
	CardInfo CardBaseResponse `json:"cardInfo"`
}
