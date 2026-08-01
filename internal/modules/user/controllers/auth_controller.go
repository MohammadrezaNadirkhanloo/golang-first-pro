package controllers

import (
	"first-app/internal/modules/user/requests/auth"
	"first-app/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c,http.StatusOK,"modules/user/html/register",gin.H{
		"title":"register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
  var json auth.RegisterRequest
    if err := c.ShouldBindJSON(&json); err != nil {
      c.Redirect(http.StatusNotFound,"/register")
      return
    }

	c.JSON(http.StatusOK,gin.H{"message":"registerdddd"})
}
