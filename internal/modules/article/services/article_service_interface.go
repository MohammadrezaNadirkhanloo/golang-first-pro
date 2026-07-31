package services

import (
	"first-app/internal/modules/article/respones"
)

type ArticleServiceInterface interface {
	GetFeaturedArticle() respones.Articles
	GetStoriesArticle() respones.Articles
	Find(id int) (respones.Article, error)
}
