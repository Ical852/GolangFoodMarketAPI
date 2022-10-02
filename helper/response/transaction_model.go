package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/transaction"
)

func ToTransactionResponse(transactionDomain domain.Transaction) transaction.TransactionResponse {
	return transaction.TransactionResponse{
		Id:          transactionDomain.Id,
		OrderId:     transactionDomain.OrderId,
		GrossAmount: transactionDomain.GrossAmount,
		PaymentUrl:  transactionDomain.PaymentUrl,
		UserId:      transactionDomain.UserId,
		ProductId:   transactionDomain.ProductId,
		Status:      transactionDomain.Status,
	}
}

func ToTransactionResponses(transactions []domain.Transaction) []transaction.TransactionResponse {
	var transactionsResponses []transaction.TransactionResponse
	for _, transactionItem := range  transactions {
		transactionsResponses = append(transactionsResponses, ToTransactionResponse(transactionItem))
	}
	return transactionsResponses
}