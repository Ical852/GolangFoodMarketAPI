package favourite

type FavouriteResponse struct {
	Id int `json:"id"`
	MarketId int `json:"market_id"`
	UserId int `json:"user_id"`
}