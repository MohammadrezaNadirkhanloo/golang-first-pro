package repositories

import (
	"first-app/internal/modules/article/models"
	"first-app/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connections(),
	}
}

func (articlerepository *ArticleRepository) List(limit int) []models.Article {
	var article []models.Article
	articlerepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&article)
	return article
}
