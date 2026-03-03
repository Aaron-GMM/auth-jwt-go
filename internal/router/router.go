package router

import "github.com/gin-gonic/gin"

func InitRouter() {
	var router *gin.Engine = gin.Default()

	initializeRouter(router)

	router.Run(":8080")

}
