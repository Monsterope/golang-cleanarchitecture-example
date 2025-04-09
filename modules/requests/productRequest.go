package requests

type (
	ItemListProd struct {
		Item []ProductCreateRequest `json:"item" validate:"required"`
	}
	ProductCreateRequest struct {
		Code       string  `json:"code" validate:"required"`
		Name       string  `json:"name" validate:"required"`
		Price      float64 `json:"price"`
		CategoryId int64   `json:"cate_id"`
	}
)
