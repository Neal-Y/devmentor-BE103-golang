package service

import (
	"errors"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/order"
	"shopping-cart/repository"
	"time"
)

type OrderService interface {
	CreateOrder(orderRequest *order.Request) (*database.Order, error)
	GetOrderByID(id int) (*database.Order, error)
	UpdateOrder(id int, orderRequest *order.Request) (*database.Order, error)
	DeleteOrder(id int) error
	ListAllOrders() ([]database.Order, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) CreateOrder(orderRequest *order.Request) (*database.Order, error) {
	totalPrice := 0.0

	for i, detail := range orderRequest.OrderDetails {
		if detail.Quantity <= 0 {
			return nil, errors.New("Quantity must be greater than zero")
		}

		product, err := s.productRepo.FindByID(detail.ProductID)
		if err != nil {
			return nil, errors.New("Product not found")
		}

		if product.Stock < detail.Quantity {
			return nil, errors.New("Insufficient stock for product " + product.Name)
		}

		if time.Now().After(product.ExpirationTime) {
			return nil, errors.New("Product " + product.Name + " is expired")
		}

		orderRequest.OrderDetails[i].Price = product.Price

		totalPrice += float64(detail.Quantity) * product.Price

		now := time.Now()
		orderRequest.OrderDetails[i].CreatedAt = now
		orderRequest.OrderDetails[i].UpdatedAt = now
	}

	order := &database.Order{
		UserID:       orderRequest.UserID,
		TotalPrice:   totalPrice,
		Note:         orderRequest.Note,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		OrderDetails: orderRequest.OrderDetails,
	}

	err := s.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	for _, detail := range order.OrderDetails {
		product, err := s.productRepo.FindByID(detail.ProductID)
		if err != nil {
			return nil, err
		}
		product.Stock -= detail.Quantity
		err = s.productRepo.Update(product)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (s *orderService) GetOrderByID(id int) (*database.Order, error) {
	return s.orderRepo.FindByID(id)
}

func (s *orderService) UpdateOrder(id int, orderRequest *order.Request) (*database.Order, error) {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("Order not found")
	}

	totalPrice := order.TotalPrice

	if len(orderRequest.OrderDetails) > 0 {
		totalPrice = 0.0
		for _, detail := range orderRequest.OrderDetails {
			product, err := s.productRepo.FindByID(detail.ProductID)
			if err != nil {
				return nil, errors.New("Product not found")
			}

			if product.Stock < detail.Quantity {
				return nil, errors.New("Insufficient stock for product " + product.Name)
			}

			if time.Now().After(product.ExpirationTime) {
				return nil, errors.New("Product " + product.Name + " is expired")
			}

			totalPrice += float64(detail.Quantity) * product.Price
		}
		order.OrderDetails = orderRequest.OrderDetails
	}

	if orderRequest.Note != "" {
		order.Note = orderRequest.Note
	}

	if orderRequest.Status != "" {
		order.Status = orderRequest.Status
	}

	order.TotalPrice = totalPrice
	order.UpdatedAt = time.Now()

	err = s.orderRepo.Update(order)
	if err != nil {
		return nil, err
	}

	for _, detail := range order.OrderDetails {
		product, err := s.productRepo.FindByID(detail.ProductID)
		if err != nil {
			return nil, err
		}
		product.Stock -= detail.Quantity
		err = s.productRepo.Update(product)
		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (s *orderService) DeleteOrder(id int) error {
	order, err := s.orderRepo.FindByID(id)
	if err != nil {
		return errors.New("Order not found")
	}

	tx := s.orderRepo.BeginTransaction()

	for _, detail := range order.OrderDetails {
		product, err := s.productRepo.FindByID(detail.ProductID)
		if err != nil {
			tx.Rollback()
			return err
		}
		product.Stock += detail.Quantity
		err = s.productRepo.Update(product)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = s.orderRepo.Delete(order)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *orderService) ListAllOrders() ([]database.Order, error) {
	return s.orderRepo.FindAll()
}
