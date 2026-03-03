package repository

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/domain/entity"
	"github.com/Aaron-GMM/auth-jwt-go/internal/handler"
	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
func (r *authRepository) CreateUser(user *entity.User) error {
	Logger := handler.Logger
	err := r.db.Create(user).Error
	if err != nil {
		Logger.ErrorF("create user error: %v", err.Error())
		return err
	}
	return nil

}
func (r *authRepository) FindUserByEmail(email string) (*entity.User, error) {
	Logger := handler.Logger
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		Logger.ErrorF("find user error: %v", err.Error())
		return nil, err
	}
	return &user, nil
}

//func (r *AuthRepository) FindByName(name string) (*entity.User, error) {
//	Logger := handler.Logger
//	var user entity.User
//	err := r.db.Where("name = ?", name).First(&user).Error
//	if err != nil {
//		Logger.ErrorF("find user error: %v", err.Error())
//		return nil, err
//	}
//	return &user, nil
//}
//func (r *AuthRepository) FindById(id uint) (*entity.User, error) {
//	Logger := handler.Logger
//	var user entity.User
//	err := r.db.Where("id = ?", id).First(&user).Error
//	if err != nil {
//		Logger.ErrorF("find user error: %v", err.Error())
//		return nil, err
//	}
//	return &user, nil
//}
