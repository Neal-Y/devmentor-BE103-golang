package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type OrderRepository interface {
	Create(order *database.Order) error
	FindByID(id int) (*database.Order, error)
	Update(order *database.Order) error
	Delete(order *database.Order) error
	FindAll() ([]database.Order, error)
	BeginTransaction() *gorm.DB
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		db: infrastructure.Db,
	}
}

func (r *orderRepository) Create(order *database.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id int) (*database.Order, error) {
	var order database.Order
	err := r.db.Preload("User").Preload("OrderDetails.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) Update(order *database.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) Delete(order *database.Order) error {
	return r.db.Delete(order).Error
}

func (r *orderRepository) FindAll() ([]database.Order, error) {
	var orders []database.Order
	err := r.db.Preload("User").Preload("OrderDetails.Product").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
