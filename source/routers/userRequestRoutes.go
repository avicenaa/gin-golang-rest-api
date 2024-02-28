package routers

import (
	"api-request/controllers"

	"github.com/gin-gonic/gin"
)

func UserRequestRoute(route *gin.Engine) {
	route.GET("/api/view_all", controllers.GetAllForms)
	route.POST("/api/user/input", controllers.CreateForm)
	route.PUT("/api/user/update/:id", controllers.UpdateFormByID)
	route.DELETE("/api/user/:id", controllers.DeleteFormByID)
	route.GET("/api/view_detail/:id", controllers.GetFormByID)
}

func CrawlerRequestRoute(route *gin.Engine) {
	route.PUT("/api/crawler/update/:id", controllers.AddCrawlerData)
}
