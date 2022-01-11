package market

type MarketResponse struct {
	Id int `json:"id"`
	CategoryId int `json:"category_id"`
	MarketName string `json:"market_name"`
	MarketImg string `json:"market_img"`
	MarketLocation string `json:"market_location"`
	MarketRating string `json:"market_rating"`
	MarketProductName string `json:"market_product_name"`
	MarketEstimate string `json:"market_estimate"`
}