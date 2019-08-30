package request

// GetCardList的Request参数
type GetCardListRequest struct {
	BoardId    string `json:"boardId"`
	CardStatus int    `json:"cardStatus"`
}

// CardList 请求入参
type CardIdList struct {
	CardIdList []string
}
