package usecase

import "github.com/kodinggo/category-service-gb1/internal/model"

type categoryUsecase struct {
	categoryRepo model.CategoryRepository
}

func NewCategoryUsecase(cr model.CategoryRepository) model.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo: cr,
	}
}

func (cu *categoryUsecase) FindCategoryByID(id string) (model.Category, error) {
	return cu.categoryRepo.FindCategoryByID(id)
}
