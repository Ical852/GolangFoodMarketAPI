package favourite

type FavouriteDeleteRequest struct {
	MarketId int `validate:"required" json:"market_id"`
	UserId int `validate:"required" json:"user_id"`
}