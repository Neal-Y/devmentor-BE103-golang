package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type AdminRepository interface {
	Create(admin *database.Admin) error
	FindByUsername(username string) (*database.Admin, error)
	FindByEmail(email string) (*database.Admin, error)
	GetAdmin() (*database.Admin, error)
	Update(admin *database.Admin) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		db: infrastructure.Db,
	}
}

func (r *adminRepository) Create(admin *database.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) FindByUsername(username string) (*database.Admin, error) {
	var admin database.Admin
	err := r.db.Where("BINARY username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) FindByEmail(email string) (*database.Admin, error) {
	var admin database.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) GetAdmin() (*database.Admin, error) {
	var admin database.Admin
	err := r.db.First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Update(admin *database.Admin) error {
	return r.db.Save(admin).Error
}
