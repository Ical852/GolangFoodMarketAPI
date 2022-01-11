package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type FavouriteRepository interface {
	Get(ctx context.Context, tx *sql.Tx, userId int) []domain.Favourite
	Check(ctx context.Context, tx *sql.Tx, favourite domain.Favourite) (domain.Favourite, error)
	Add(ctx context.Context, tx *sql.Tx, favourite domain.Favourite) domain.Favourite
	Delete(ctx context.Context, tx *sql.Tx, favourite domain.Favourite)
}