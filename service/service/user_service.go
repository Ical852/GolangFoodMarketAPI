package service

import (
	"Go-FoodMarket/model/web/user"
	"context"
)

type UserService interface {
	Login(ctx context.Context, request user.UserLoginRequest) user.UserResponse
	Register(ctx context.Context, request user.UserRegisterRequest) user.UserResponse
	GetUser(ctx context.Context, userId int) user.UserResponse
	Update(ctx context.Context, request user.UserUpdateRequest) user.UserResponse
}