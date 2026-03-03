package database

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/domain/entity"

	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(dsn string) (*gorm.DB, error) {
	logger := config.GetLogger("postegresql")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error connecting to database: %s", err.Error())
		return nil, err // Pare aqui se der erro!
	}

	// Só migra se a conexão for bem-sucedida
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.ErrorF("Error migrating database: %s", err.Error())
		return nil, err
	}

	return db, nil

}
