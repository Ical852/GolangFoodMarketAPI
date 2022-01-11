package response

import (
	"Go-FoodMarket/model/domain"
	"Go-FoodMarket/model/web/user"
)

func ToUserResponse(userDomain domain.User) user.UserResponse {
	return user.UserResponse{
		Id:          userDomain.Id,
		FullName:    userDomain.FullName,
		Username:    userDomain.Username,
		Password:    userDomain.Password,
		Email:       userDomain.Email,
		UserImg:     userDomain.UserImg,
		PhoneNumber: userDomain.PhoneNumber,
		RoleId:      userDomain.RoleId,
		UserStatus:  userDomain.UserStatus,
	}
}

func ToUserResponses(users []domain.User) []user.UserResponse {
	var userResponses []user.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}