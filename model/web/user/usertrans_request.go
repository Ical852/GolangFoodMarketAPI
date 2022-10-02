package user

type UserTransRequest struct {
	FullName string `validate:"required" json:"full_name"`
	Email string `validate:"required" json:"email"`
	Phone string `validate:"required" json:"phone"`
}
