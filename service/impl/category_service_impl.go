package impl

import (
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/web/category"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) service.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Get(ctx context.Context) []category.CategoryReponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.Get(ctx, tx)
	return response.ToCategoryResponses(categories)
}
