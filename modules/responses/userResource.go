package responses

import (
	"cleanarchitecture-example/modules/entities"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SafeCustomer struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func ModelUser(u *entities.User) User {
	return User{
		Id:        u.Id,
		Username:  u.Username,
		Name:      u.Name,
		UserType:  u.UserType,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.CreatedAt,
	}
}

func SafeModelCustomer(u *entities.User) SafeCustomer {
	return SafeCustomer{
		Username:  u.Username,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}
}
