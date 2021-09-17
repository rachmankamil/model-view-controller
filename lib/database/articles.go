package database

import (
	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/models"
)

func GetArticles() (*[]models.Article, error) {
	var articles []models.Article

	if err := config.DB.Find(&articles).Error; err != nil {
		return &[]models.Article{}, err
	}

	return &articles, nil
}

func GetbyIDArticle(id int) (*models.Article, error) {
	var article models.Article

	if err := config.DB.Where("id = ?", id).First(&article).Error; err != nil {
		return &models.Article{}, err
	}

	return &article, nil
}

func StoreArticle(article models.Article) (*models.Article, error) {
	if err := config.DB.Save(&article).Error; err != nil {
		return &models.Article{}, err
	}

	return &article, nil
}
