package database

import (
	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/models"
)

func GetArticles(title string, rowPerPage int, page int) (*[]models.Article, error) {
	if rowPerPage == 0 {
		rowPerPage = 15
	}
	if page == 0 {
		page = 1
	}

	var articles []models.Article

	query := config.DB

	if title != "" {
		query = query.Where("title LIKE '%?%'", title)
	}

	query.Limit(rowPerPage)
	query.Offset((rowPerPage * page) - rowPerPage)

	if err := query.Find(&articles).Error; err != nil {
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

func Update(article models.Article) (*models.Article, error) {
	if err := config.DB.Updates(&article).Error; err != nil {
		return &models.Article{}, err
	}

	return &article, nil
}
