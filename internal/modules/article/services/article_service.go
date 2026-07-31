package services

import (
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
