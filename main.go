package main

import (
	"log"

	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/config"
	"github.com/Aaron-GMM/auth-jwt-go/internal/infra/database"
	"github.com/Aaron-GMM/auth-jwt-go/internal/router"
)

func main() {
	// 1. Config & Logger
	cfg := config.LoadConfig()
	logger := config.GetLogger("main")

	// 2. Database (Usando a URL do .env)
	db, err := database.InitializeDatabase(cfg.DBURL)
	if err != nil {
		log.Fatalf("Nao foi possivel conectar ao banco: %v", err)
	}

	logger.InforF("Serviço de Autenticação pronto!")

	// O próximo passo será ligar o Router do Gin aqui
	router.InitRouter(db)
}
