package transaction

type TransactionCreateRequest struct {
	OrderId int `validate:"required" json:"order_id"`
	GrossAmount int `validate:"required" json:"gross_amount"`
	UserId int `validate:"required" json:"user_id"`
	ProductId int `validate:"required" json:"product_id"`
	Status string `validate:"required" json:"status"`
	FullName string `validate:"required" json:"full_name"`
	Email string `validate:"required" json:"email"`
	Phone string `validate:"required" json:"phone"`
}