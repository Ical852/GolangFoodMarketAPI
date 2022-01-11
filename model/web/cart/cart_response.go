package cart

type CartResponse struct {
	Id int `json:"id"`
	ProductId int `json:"product_id"`
	UserId int `json:"user_id"`
	Qty int `json:"qty"`
	Status string `json:"status"`
}