package main

import (
	"log"
	"net"

	grpcSvc "github.com/kodinggo/category-service-gb1/internal/delivery/grpc"
	"github.com/kodinggo/category-service-gb1/internal/repository"
	"github.com/kodinggo/category-service-gb1/internal/usecase"
	pb "github.com/kodinggo/category-service-gb1/pb/category"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

	categoryRepo := repository.NewCategoryRepository()
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	categoryService := grpcSvc.NewCategoryService(categoryUsecase)

	pb.RegisterCategoryServiceServer(s, categoryService)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server at :3200")

	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
