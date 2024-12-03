package services

import (
	"Go-GitHub-Projects/Banking-App/models"
	"Go-GitHub-Projects/Banking-App/storage"
	"Go-GitHub-Projects/Banking-App/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	RegisterUser(userName string, password string) error
	AuthenticateUser(user *models.AccountHolder) (string, error)
}

type UserService struct {
	userRepo storage.UserRepository
}

func NewUserService(userRepo storage.UserRepository) *UserService {

	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) RegisterUser(userName, password string) error {

	existingUser, err := s.userRepo.GetUserByUserName(userName)

	var user models.AccountHolder

	if err != nil && existingUser == user {

		return fmt.Errorf("user already exists")
	}

	return s.userRepo.SaveUser(userName, password)
}

func (s *UserService) AuthenticateUser(userName, password string) (string, error) {

	user, err := s.userRepo.GetUserByUserName(userName)

	if err != nil {

		return "", fmt.Errorf("invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {

		return "", fmt.Errorf("invalid credentials")
	}

	token, err := utils.GenerateJwtToken(userName)

	if err != nil {

		return "", err
	}
	return token, nil
}
