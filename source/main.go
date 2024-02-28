package main

import (
	_ "api-request/docs"
	"api-request/routers"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routers.UserRequestRoute(router)
	routers.CrawlerRequestRoute(router)

	router.Run("localhost:8080")
}
