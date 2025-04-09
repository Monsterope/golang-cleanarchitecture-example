package requests

type (
	CategoryCreateRequest struct {
		Item []CategoryCreateItemRequest `json:"item" validate:"required,dive"`
	}
	CategoryCreateItemRequest struct {
		Name string `json:"name" validate:"required"`
	}
	CategoryUpdateRequest struct {
		Name string `json:"name" validate:"required"`
	}
)

func (CategoryCreateItemRequest) TableName() string {
	return "categories"
}
