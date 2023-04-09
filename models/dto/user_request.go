package dto

type UsersCreateRequest struct {
	Username string `validate:"required,min=1,max=100" json:"username"`
	Email    string `validate:"required,min=1,max=50,email" json:"email"`
	Password string `validate:"required,min=6" json:"password"`
}

type UsersUpdateRequest struct {
	Id        string `validate:"required" json:"id"`
	Username  string `validate:"required,min=1,max=100" json:"username"`
	UpdatedAt int64  `json:"updated_at"`
}

type UserAuthRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
