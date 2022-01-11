package user

type UserResponse struct {
	Id int `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	UserImg string `json:"user_img"`
	PhoneNumber string `json:"phone_number"`
	RoleId int `json:"role_id"`
	UserStatus int `json:"user_status"`
}