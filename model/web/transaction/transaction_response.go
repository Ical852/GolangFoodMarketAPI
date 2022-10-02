package transaction

type TransactionResponse struct {
	Id int `json:"id"`
	OrderId int `json:"order_id"`
	GrossAmount int `json:"gross_amount"`
	PaymentUrl string `json:"payment_url"`
	UserId int `json:"user_id"`
	ProductId int `json:"product_id"`
	Status string `json:"status"`
}
