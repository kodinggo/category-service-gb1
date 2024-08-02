package model

type Category struct {
	Id        string `json:"id"`
	Name      string `json:"story_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryUsecase interface {
	FindCategoryByID(id string) (Category, error)
}

type CategoryRepository interface {
	FindCategoryByID(id string) (Category, error)
}
