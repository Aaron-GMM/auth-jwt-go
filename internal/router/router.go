package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	var router *gin.Engine = gin.Default()

	initializeRouter(router, db)

	router.Run(":8080")

}
