package routes

import (
	Controller "first-app/internal/modules/article/controllers"
	homeCtrl "first-app/internal/modules/home/controllers"
	"first-app/internal/modules/user/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	articleControllers := Controller.New()
	userControllers := controllers.New()
	router.GET("/", homeController.Index)
	router.GET("/article/:id", articleControllers.Show)
	router.GET("/register", userControllers.Register)
	router.POST("/register", userControllers.HandleRegister)
}
