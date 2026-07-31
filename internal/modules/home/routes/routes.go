package routes

import (
	Controller "first-app/internal/modules/article/controllers"
	homeCtrl "first-app/internal/modules/home/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	articleControllers := Controller.New()
	router.GET("/", homeController.Index)
	router.GET("/article/:id", articleControllers.Show)
}
