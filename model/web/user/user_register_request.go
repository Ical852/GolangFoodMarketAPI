package user

type UserRegisterRequest struct {
	FullName string `validate:"required" json:"full_name"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Email string `validate:"required" json:"email"`
	UserImg string `validate:"required" json:"user_img"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	RoleId int `validate:"required" json:"role_id"`
	UserStatus int `validate:"required" json:"user_status"`
}