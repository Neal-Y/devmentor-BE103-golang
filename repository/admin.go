package repository

import (
	"gorm.io/gorm"
	"shopping-cart/infrastructure"
	"shopping-cart/model/database"
)

type AdminRepository interface {
	Create(admin *database.Admin) error
	FindByUsername(username string) (*database.Admin, error)
	FindAll() ([]database.Admin, error)
	FindByID(id int) (*database.Admin, error)
	Update(admin *database.Admin) error
	Delete(admin *database.Admin) error
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
	err := r.db.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) FindAll() ([]database.Admin, error) {
	var admins []database.Admin
	err := r.db.Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil

}

func (r *adminRepository) FindByID(id int) (*database.Admin, error) {
	var admin database.Admin
	err := r.db.First(&admin, id).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) Update(admin *database.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) Delete(admin *database.Admin) error {
	return r.db.Delete(admin).Error
}
