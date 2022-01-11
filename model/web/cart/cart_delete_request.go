package cart

type CartDeleteRequest struct {
	ProductId int `validate:"required" json:"product_id"`
	UserId int `validate:"required" json:"user_id"`
}