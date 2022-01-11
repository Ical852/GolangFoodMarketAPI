package repository

import (
	"Go-FoodMarket/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Get(ctx context.Context, tx *sql.Tx) []domain.Category
}