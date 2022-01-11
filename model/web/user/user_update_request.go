package user

type UserUpdateRequest struct {
	Id int `validate:"required" json:"id"`
	FullName string `validate:"required" json:"full_name"`
	Password string `validate:"required" json:"password"`
	PhoneNumber string `validate:"required" json:"phone_number"`
}