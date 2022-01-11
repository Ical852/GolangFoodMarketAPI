package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/category"
)

func ToCategoryResponse(categoryDomain domain.Category) category.CategoryReponse {
	return category.CategoryReponse{
		Id:           categoryDomain.Id,
		CategoryName: categoryDomain.CategoryName,
	}
}

func ToCategoryResponses(categories []domain.Category) []category.CategoryReponse {
	var categoryResponses []category.CategoryReponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}