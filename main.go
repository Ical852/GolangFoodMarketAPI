package main

import (
	"Go-FoodMarket/app"
	impl3 "Go-FoodMarket/controller/impl"
	"Go-FoodMarket/exception"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/middleware"
	"Go-FoodMarket/repository/impl"
	impl2 "Go-FoodMarket/service/impl"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	userRepository := impl.NewUserRepository()
	userService := impl2.NewUserService(userRepository, db, validate)
	userController := impl3.NewUserController(userService)

	productsRepository := impl.NewProductRepository()
	productsService := impl2.NewProductsService(productsRepository, db, validate)
	productsController := impl3.NewProductsController(productsService)

	marketRepository := impl.NewMarketRepository()
	marketService := impl2.NewMarketService(marketRepository, db, validate)
	marketController := impl3.NewMarketController(marketService)

	favouriteRepository := impl.NewFavouriteRepositoy()
	favouriteService := impl2.NewFavouriteService(favouriteRepository, db, validate)
	favouriteController := impl3.NewFavouriteController(favouriteService)

	categoryRepository := impl.NewCategoryRepository()
	categoryService := impl2.NewCategoryService(categoryRepository, db, validate)
	categoryController := impl3.NewCategoryController(categoryService)

	cartRepository := impl.NewCartRepository()
	cartService := impl2.NewCartService(cartRepository, db, validate)
	cartController := impl3.NewCartController(cartService)

	router := httprouter.New()

	router.POST("/api/user/login", userController.Login)
	router.POST("/api/user/register", userController.Register)
	router.GET("/api/user/get/:userId", userController.GetUser)
	router.PUT("/api/user/update/:userId", userController.Update)

	router.GET("/api/products/get/:marketId", productsController.Get)
	router.GET("/api/products/detail/:productId", productsController.GetById)

	router.GET("/api/market/get", marketController.Get)
	router.GET("/api/market/getrand", marketController.GetRandom)
	router.GET("/api/market/detail/:marketId", marketController.GetById)
	router.GET("/api/market/category/:categoryId", marketController.GetByCategory)
	router.GET("/api/market/search/:searchValue", marketController.GetBySearch)

	router.GET("/api/favourite/get/:userId", favouriteController.Get)
	router.POST("/api/favourite/check", favouriteController.Check)
	router.POST("/api/favourite/add", favouriteController.Add)
	router.POST("/api/favourite/delete", favouriteController.Delete)

	router.GET("/api/category/get", categoryController.Get)

	router.GET("/api/cart/get/:userId", cartController.Get)
	router.POST("/api/cart/add", cartController.Add)
	router.POST("/api/cart/delete", cartController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}