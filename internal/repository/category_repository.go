package repository

import "kodinggo/category-service-gb1/internal/model"

type categoryRepository struct {
}

func NewCategoryRepository() model.CategoryRepository {
	return &categoryRepository{}
}

func (cr *categoryRepository) FindCategoryByID(id string) (model.Category, error) {
	var category model.Category

	// select * from categorys where id = 1
	if id != "1" {
		return category, nil
	}

	return model.Category{
		Id:        "1",
		Name:      "News",
		CreatedAt: "2021-01-01T00:00:00Z",
		UpdatedAt: "2021-01-01T00:00:00Z",
	}, nil
}
