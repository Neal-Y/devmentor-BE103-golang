package database

import (
	"time"
)

type Order struct {
	ID           int           `json:"id" gorm:"primary_key;autoIncrement"`
	UserID       int           `json:"user_id" gorm:"not null"`
	TotalPrice   float64       `json:"total_price" gorm:"type:decimal(10,2)"`
	Note         string        `json:"note" gorm:"type:text"`
	Status       string        `json:"status" gorm:"type:enum('pending', 'completed', 'cancelled');default:'pending'"`
	CreatedAt    time.Time     `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time     `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	User         *User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderDetails []OrderDetail `json:"order_details,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderDetail struct {
	ID        int       `json:"id" gorm:"primary_key;autoIncrement"`
	OrderID   int       `json:"order_id" gorm:"not null"`
	ProductID int       `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Order     *Order    `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product   *Product  `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (Order) TableName() string {
	return "orders"
}

func (OrderDetail) TableName() string {
	return "order_details"
}
