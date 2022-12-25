package services

import (
	"context"
	"web1/internal/core"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*core.Product, error)
	FindById(ctx context.Context, id string) (*core.Product, error)
	Save(ctx context.Context, product *core.Product) (*core.Product, error)
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (service *ProductService) GetAll(ctx context.Context) ([]*core.Product, error) {
	return service.productRepository.FindAll(ctx)
}

func (service *ProductService) GetById(ctx context.Context, id string) (*core.Product, error) {
	return service.productRepository.FindById(ctx, id)
}

func (service *ProductService) AddProduct(ctx context.Context, product *core.Product) (*core.Product, error) {
	return service.productRepository.Save(ctx, product)
}