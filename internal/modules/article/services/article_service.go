package services

import (
	"errors"
	"first-app/internal/modules/article/repositories"
	"first-app/internal/modules/article/respones"
)

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: repositories.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticle() respones.Articles {
	articles := articleService.articleRepository.List(4)

	return respones.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticle() respones.Articles {
	articles := articleService.articleRepository.List(6)
	return respones.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (respones.Article, error) {
	var response respones.Article

	article := articleService.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("error")
	}
	return respones.ToArticle(article), nil
}
