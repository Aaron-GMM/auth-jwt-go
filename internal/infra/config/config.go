package config

import (
	"log"
	"os"

	"github.com/Aaron-GMM/auth-jwt-go/internal/infra"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	DBURL  string
	JWTKey string
}

var (
	logger *infra.Logger
	db     *gorm.DB
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		// Change log.Fatal to log.Println so it doesn't crash the container
		log.Println("Aviso: arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

	return &Config{
		DBURL:  os.Getenv("DB_URL"),
		JWTKey: os.Getenv("SECRET_KEY_JWT"),
	}
}
func GetDb() *gorm.DB {
	return db
}
func GetLogger(p string) *infra.Logger {
	if logger == nil {
	}
	logger = infra.NewLogger(p)
	return logger
}
