package response

// Card 基础数据
type CardBaseResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	CardStatus int    `json:"cardStatus"`
}

// GetCardList列表返回数据
type CardResponse struct {
	CardInfo CardBaseResponse `json:"cardInfo"`
}

type ListResponse struct {
	List  []CardResponse `json:"list"`
	Count int            `json:"count"`
}
