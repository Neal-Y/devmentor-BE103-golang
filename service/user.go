package service

import (
	"fmt"
	"shopping-cart/builder"
	"shopping-cart/config"
	"shopping-cart/constant"
	"shopping-cart/model/database"
	"shopping-cart/model/datatransfer/user"
	"shopping-cart/repository"
	"shopping-cart/util"
)

type UserService interface {
	CreateUser(user *database.User) error
	GetUserByID(id int) (*database.User, error)
	UpdateUser(user *database.User) error
	DeleteUser(user *database.User) error
	FindByLineID(lineID string) (*database.User, error)
	SaveOrUpdateUser(user *database.User) error
	ExchangeTokenAndGetProfile(code string) (*database.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *database.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserByID(id int) (*database.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *database.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(user *database.User) error {
	return s.repo.Delete(user)
}

func (s *userService) FindByLineID(lineID string) (*database.User, error) {
	return s.repo.FindByLineID(lineID)
}

func (s *userService) SaveOrUpdateUser(user *database.User) error {
	existingUser, err := s.repo.FindByLineID(user.LineID)
	if err != nil {
		return s.repo.Create(user)
	}
	user.ID = existingUser.ID
	return s.repo.Update(user)
}

func (s *userService) ExchangeTokenAndGetProfile(code string) (*database.User, error) {
	var tokenData *user.LineTokenResponse
	httpBuilder := builder.NewHttpClient[*user.LineTokenResponse]()

	err := httpBuilder.
		WithMethodPost().
		WithURL(constant.LineTokenURL).
		WithFormData("grant_type", "authorization_code").
		WithFormData("code", code).
		WithFormData("redirect_uri", config.AppConfig.LineRedirectURI).
		WithFormData("client_id", config.AppConfig.LineClientID).
		WithFormData("client_secret", config.AppConfig.LineClientSecret).
		UserHeaderFormUrlencoded().
		Build(&tokenData)
	if err != nil {
		return nil, err
	}

	if tokenData.AccessToken == "" {
		return nil, fmt.Errorf("failed to parse access token")
	}

	var profileData *user.LineProfileResponse
	profileData, err = util.ParseIDToken(tokenData.IDToken)
	if err != nil {
		return nil, err
	}

	user := builder.NewUserBuilder().
		WithLineID(profileData.UserID).
		WithDisplayName(profileData.DisplayName).
		WithEmail(profileData.Email).
		WithLineToken(tokenData.AccessToken).
		Build()

	return user, nil
}
