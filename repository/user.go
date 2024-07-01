package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type UserRepository interface {
	Create(user *database.User) error
	FindByID(id int) (*database.User, error)
	Update(user *database.User) error
	Delete(user *database.User) error
	FindByLineID(lineID string) (*database.User, error)
	DeleteTx(tx *gorm.DB, id int) error
	BeginTransaction() *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: infrastructure.Db,
	}
}

func (r *userRepository) Create(user *database.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id int) (*database.User, error) {
	var user database.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *database.User) error {
	return r.db.Updates(user).Error
}

func (r *userRepository) Delete(user *database.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) FindByLineID(lineID string) (*database.User, error) {
	var user database.User
	err := r.db.Where("line_id = ?", lineID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) DeleteTx(tx *gorm.DB, id int) error {
	var user database.User
	err := tx.First(&user, id).Error
	if err != nil {
		return err
	}
	return tx.Delete(&user).Error
}

func (r *userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
