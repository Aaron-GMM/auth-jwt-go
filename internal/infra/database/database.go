package database

import (
	"github.com/Aaron-GMM/auth-jwt-go/internal/domain/entity"

	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitializeDatabase(dns string) (*gorm.DB, error) {
	logger := config.GetLogger("postegresql")
	var err error
	Db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		logger.ErrorF("Error connecting to database: %v", err)
	}
	err = Db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.ErrorF("Error auto-migrating database: %v", err)
	}
	return Db, err

}
