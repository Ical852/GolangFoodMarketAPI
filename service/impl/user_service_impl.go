package impl

import (
	"Go-FoodMarket/exception"
	"Go-FoodMarket/helper"
	"Go-FoodMarket/helper/response"
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/user"
	"Go-FoodMarket/repository/repository"
	"Go-FoodMarket/service/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) service.UserService  {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, request user.UserLoginRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err = service.UserRepository.Login(ctx, tx, user)
	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) Register(ctx context.Context, request user.UserRegisterRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		FullName:    request.FullName,
		Username:    request.Username,
		Password:    request.Password,
		Email:       request.Email,
		UserImg:     request.UserImg,
		PhoneNumber: request.PhoneNumber,
		RoleId:      request.RoleId,
		UserStatus:  request.UserStatus,
	}

	user = service.UserRepository.Register(ctx, tx, user)
	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) GetUser(ctx context.Context, userId int) user.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUser(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return response.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetUser(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.FullName = request.FullName
	user.Password = request.Password
	user.PhoneNumber = request.PhoneNumber

	user = service.UserRepository.Update(ctx, tx, user)
	return response.ToUserResponse(user)
}
