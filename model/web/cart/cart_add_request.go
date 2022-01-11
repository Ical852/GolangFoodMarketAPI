package cart

type CartAddRequest struct {
	ProductId int `validate:"required" json:"product_id"`
	UserId int `validate:"required" json:"user_id"`
	Qty int `validate:"required" json:"qty"`
	Status string `validate:"required" json:"status"`
}