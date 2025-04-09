package responses

import (
	"cleanarchitecture-example/modules/entities"
	"time"
)

type CategoryResource struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetCategoryResource(c *entities.Category) CategoryResource {
	return CategoryResource{
		Id:        c.Id,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
