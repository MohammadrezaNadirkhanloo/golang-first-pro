package controllers

import (
	"first-app/internal/modules/article/services"
	"first-app/pkg/html"
	"net/http"

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

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home page",
		"Featured": controller.articleService.GetFeaturedArticle(),
		"stories":  controller.articleService.GetStoriesArticle(),
	})

}
