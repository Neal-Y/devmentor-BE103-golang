package service

import (
	"errors"
	"shopping-cart/builder"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/product"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type ProductService interface {
	UpdateProduct(id int, productDto *product.Update) error
	CreateProduct(productDto *product.Payload) (*database.Product, error)
	DeleteProduct(id int) error
	FindByID(id int) (*database.Product, error)
	SearchProducts(params util.SearchContainer) ([]database.ProductWithTime, int64, error)
	GetByID(id int) (*database.ProductWithTime, error)
}

type productService struct {
	productRepo       repository.ProductRepository
	notificationCache *util.NotificationCache
}

func NewProductService(repo repository.ProductRepository, notificationCache *util.NotificationCache) ProductService {
	return &productService{
		productRepo:       repo,
		notificationCache: notificationCache,
	}
}

func (s *productService) UpdateProduct(id int, productDto *product.Update) error {
	product, err := s.productRepo.InternalFindByID(id)
	if err != nil {
		return err
	}

	var number int

	if productDto.Stock != 0 {
		number = product.Stock + productDto.Stock
		s.notificationCache.Set(product.ID, number)
	}

	product = builder.NewProductBuilder().
		SetID(product.ID).
		SetName(productDto.Name).
		SetPicture(productDto.Picture).
		SetPrice(productDto.Price).
		SetStock(number).
		SetDescription(productDto.Description).
		SetExpirationTime(productDto.ExpirationTime).
		SetSupplier(productDto.Supplier).
		Build()

	return s.productRepo.Update(product)
}

func (s *productService) CreateProduct(productDto *product.Payload) (*database.Product, error) {
	var check database.Product

	err := s.productRepo.FindByName(productDto.Name, &check)
	if err == nil {
		return nil, errors.New("product name already exists")
	}

	product := builder.NewProductBuilder().
		SetName(productDto.Name).
		SetPicture(productDto.Picture).
		SetPrice(productDto.Price).
		SetStock(productDto.Stock).
		SetDescription(productDto.Description).
		SetExpirationTime(productDto.ExpirationTime).
		SetSupplier(productDto.Supplier).
		Build()

	err = s.productRepo.Create(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProduct(id int) error {
	product, err := s.productRepo.InternalFindByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.SoftDelete(product)
}

func (s *productService) FindByID(id int) (*database.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) SearchProducts(params util.SearchContainer) ([]database.ProductWithTime, int64, error) {
	return s.productRepo.SearchProducts(params.Keyword, params.StartDate, params.EndDate, params.Offset, params.Limit)
}

func (s *productService) GetByID(id int) (*database.ProductWithTime, error) {
	return s.productRepo.FindByIDAdmin(id)
}
