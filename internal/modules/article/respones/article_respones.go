package respones

import (
	"first-app/internal/modules/article/models"
	"first-app/internal/modules/user/respones"
	"fmt"
)

type Article struct {
	ID        uint
	Image     string
	Title     string
	Content   string
	CreatedAt string
	User      respones.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article models.Article) Article {
	return Article{
		ID:        article.ID,
		Image:     "/assets/img/demopic/10.jpg",
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      respones.ToUser(article.User),
	}
}

func ToArticles(articles []models.Article) Articles {
	var response Articles

	for _, item := range articles {
		response.Data = append(response.Data, ToArticle(item))
	}

	return response
}
