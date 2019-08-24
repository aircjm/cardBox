package request

// GetCardList的Request参数
type GetCardListRequest struct {
	HaveAnki int    `json:"haveAnki"` // 0 表示无条件默认所有数据 1 表示已经有anki数据了 2表示还未生产anki note
	BoardId  string `json:"boardId"`
}
