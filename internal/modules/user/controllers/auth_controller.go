package controllers

import (
	"first-app/internal/modules/user/requests/auth"
	"first-app/internal/modules/user/services"
	"first-app/pkg/html"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService services.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: services.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var json auth.RegisterRequest
	if err := c.ShouldBind(&json); err != nil {
		c.Redirect(http.StatusNotFound, "/register")
		return
	}

	user, err := controller.userService.Creat(json)
	if err != nil {
		c.Redirect(http.StatusNotFound, "/register")
		return
	}

	log.Printf("create success - %s", user.Name)
	c.Redirect(http.StatusFound, "/")
}
