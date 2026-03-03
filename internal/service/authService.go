package service

import (
	"errors"

	"github.com/Aaron-GMM/auth-jwt-go/internal/domain/entity"
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler"
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler/requestDTOs"
	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/config"
	"github.com/Aaron-GMM/auth-jwt-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var Logger = handler.Logger

type AuthService interface {
	Register(req requestDTOs.RegisterRequestDTO) (string, error)
	Login(dto requestDTOs.LoginRequestDTO) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}
func (service *authService) Register(userReq requestDTOs.RegisterRequestDTO) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		Logger.ErrorF("Error generating password hash: %v", err.Error())
		return "", err
	}
	userRegister := &entity.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
	}

	err = service.repo.CreateUser(userRegister)
	if err != nil {
		Logger.ErrorF("Error creating user: %v", err.Error())
		return " ", err
	}
	token, err := config.GenerateToken(uint(userRegister.ID))
	if err != nil {
		Logger.ErrorF("Error generating token: %v", err.Error())
		return "", errors.New("error generating token")
	}
	return token, nil
}
func (service *authService) Login(req requestDTOs.LoginRequestDTO) (string, error) {
	user, err := service.repo.FindUserByEmail(req.Email)
	if err != nil {
		Logger.ErrorF("Error finding user: %v", err.Error())
		return "", errors.New("user not found")
	}
	isValid := config.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		Logger.ErrorF("Error finding user: %v", req.Email)
		return "", errors.New("invalid password")
	}
	token, err := config.GenerateToken(uint(user.ID))
	if err != nil {
		Logger.ErrorF("Error generating token: %v", err.Error())
		return "", errors.New("error generating token")
	}
	return token, nil
}
