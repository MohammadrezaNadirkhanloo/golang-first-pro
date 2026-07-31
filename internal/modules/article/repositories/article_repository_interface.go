package repositories

import "first-app/internal/modules/article/models"

type ArticleRepositoryInterface interface{
	List(limit int) []models.Article
}