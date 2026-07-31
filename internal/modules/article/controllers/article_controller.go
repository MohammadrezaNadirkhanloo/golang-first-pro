package Controller

import (
	"first-app/internal/modules/article/services"
	"first-app/pkg/html"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService services.ArticleService
}

func New() *Controller {
	return &Controller{
		articleService: *services.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"message": "error"})
		return
	}

	html.Render(c, http.StatusInternalServerError, "templates/errors/html/show", gin.H{"article": article})
}
