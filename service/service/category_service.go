package service

import (
	"Go-FoodMarket/model/web/category"
	"context"
)

type CategoryService interface {
	Get(ctx context.Context) []category.CategoryReponse
}