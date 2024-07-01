package service

import (
	"errors"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/product"
	"shopping-cart/repository"
	"time"
)

type ProductService interface {
	UpdateProduct(id int, productDto *product.Update) (*database.Product, error)
	CreateProduct(productDto *product.Payload) (*database.Product, error)
	DeleteProduct(id int) error
	FindAllProducts() ([]database.Product, error)
	FindByID(id int) (*database.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: repo,
	}
}

func (s *productService) UpdateProduct(id int, productDto *product.Update) (*database.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = productDto.Name
	product.Picture = productDto.Picture
	product.Price = productDto.Price
	product.Stock = productDto.Stock
	product.Description = productDto.Description
	product.ExpirationTime = productDto.ExpirationTime

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) CreateProduct(productDto *product.Payload) (*database.Product, error) {
	var check database.Product
	err := s.productRepo.FindByName(productDto.Name, &check)
	if err == nil {
		return nil, errors.New("product name already exists")
	}

	product := &database.Product{
		Name:           productDto.Name,
		Picture:        productDto.Picture,
		Price:          productDto.Price,
		Stock:          productDto.Stock,
		Description:    productDto.Description,
		ExpirationTime: productDto.ExpirationTime,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = s.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id int) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(product)
}

func (s *productService) FindAllProducts() ([]database.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) FindByID(id int) (*database.Product, error) {
	return s.productRepo.FindByID(id)
}
