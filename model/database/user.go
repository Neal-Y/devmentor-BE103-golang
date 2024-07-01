package database

import (
	"time"
)

type User struct {
	ID          int    `gorm:"primary_key"`
	LineID      string `gorm:"not null"`
	DisplayName string `gorm:"not null"`
	Email       string
	LineToken   string
	Phone       string    `gorm:"type:varchar(15)"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	IsMember    bool      `gorm:"default:false"`
	SoftDeleteModel
}

func (User) TableName() string {
	return "users"
}
