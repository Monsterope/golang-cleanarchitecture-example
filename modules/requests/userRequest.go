package requests

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	CreateUserRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
	}

	UpdateUserRequest struct {
		Name   string `json:"name" validate:"required"`
		Status int    `json:"status"`
	}
)

func (CreateUserRequest) TableName() string {
	return "users"
}

func (UpdateUserRequest) TableName() string {
	return "users"
}

func (LoginRequest) TableName() string {
	return "users"
}
