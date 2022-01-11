package products

type ProductsResponse struct {
	Id int `json:"id"`
	MarketId int `json:"market_id"`
	ProductName string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	ProductImg string `json:"product_img"`
	ProductStock string `json:"product_stock"`
	ProductStatus string `json:"product_status"`
}