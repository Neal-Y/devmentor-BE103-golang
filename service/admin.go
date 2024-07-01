package service

import (
	"errors"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/admin"
	"shopping-cart/repository"
	"shopping-cart/util"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	RegisterAdmin(admin *admin.Request) error
	Login(username, password string) (string, error)
	GetAdminByID(id int) (*database.Admin, error)
	GetAllAdmin() ([]database.Admin, error)
	UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error)
	DeleteAdmin(id int) error
}

type adminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return &adminService{
		adminRepo: adminRepo,
	}
}

func (s *adminService) RegisterAdmin(req *admin.Request) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin := &database.Admin{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}
	return s.adminRepo.Create(admin)
}

func (s *adminService) Login(username, password string) (string, error) {
	admin, err := s.adminRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := util.GenerateJWT(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) GetAllAdmin() ([]database.Admin, error) {
	return s.adminRepo.FindAll()
}

func (s *adminService) GetAdminByID(id int) (*database.Admin, error) {
	return s.adminRepo.FindByID(id)
}

func (s *adminService) UpdateAdmin(id int, req *admin.UpdateRequest) (*database.Admin, error) {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		admin.Username = req.Username
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		admin.PasswordHash = string(hashedPassword)
	}

	err = s.adminRepo.Update(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) DeleteAdmin(id int) error {
	admin, err := s.adminRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.adminRepo.Delete(admin)
}
