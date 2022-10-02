package domain

type Transaction struct {
	Id int
	OrderId int
	GrossAmount int
	PaymentUrl string
	UserId int
	ProductId int
	Status string
}