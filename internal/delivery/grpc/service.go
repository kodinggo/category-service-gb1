package grpc

import (
	"github.com/kodinggo/category-service-gb1/internal/model"
	pb "github.com/kodinggo/category-service-gb1/pb/category"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	categoryUsecase model.CategoryUsecase
}

func NewCategoryService(categoryUsecase model.CategoryUsecase) *CategoryService {
	return &CategoryService{
		categoryUsecase: categoryUsecase,
	}
}
