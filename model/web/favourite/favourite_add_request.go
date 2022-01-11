package favourite

type FavouriteAddRequest struct {
	MarketId int `validate:"required" json:"market_id"`
	UserId int `validate:"required" json:"user_id"`
}