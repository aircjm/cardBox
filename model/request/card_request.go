package request

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PageSize    int `json:"pageSize"`
}

// GetCardList的Request参数
type GetCardListRequest struct {
	BoardId    string `json:"boardId"`
	CardStatus *int   `json:"cardStatus, omitempty"`
	Pagination Pagination
}

// CardList 请求入参
type CardIdList struct {
	CardIdList []string
}
