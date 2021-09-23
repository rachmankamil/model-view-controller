package controllers

import (
	"time"

	"github.com/rachmankamil/kampus-merdeka-b/models"
)

type RequestArticle struct {
	Title   string `form:"title"`
	Author  string `form:"author"`
	RawHTML string `form:"raw_html"`
}

func (req *RequestArticle) toModel() models.Article {
	return models.Article{
		Title:   req.Title,
		Author:  req.Author,
		RawHTML: req.RawHTML,
	}
}

type ResponseArticle struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	RawHTML   string    `json:"raw_html"`
	Path      string    `json:"path_thumb"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newResponse(mdlArticles models.Article) ResponseArticle {
	return ResponseArticle{
		ID:        mdlArticles.Model.ID,
		Title:     mdlArticles.Title,
		Author:    mdlArticles.Author,
		RawHTML:   mdlArticles.RawHTML,
		Path:      mdlArticles.ThumbnailPath,
		CreatedAt: mdlArticles.CreatedAt,
		UpdatedAt: mdlArticles.UpdatedAt,
	}
}

func newResponseArray(mdlArticles []models.Article) []ResponseArticle {
	var resp []ResponseArticle

	for _, value := range mdlArticles {
		resp = append(resp, newResponse(value))
	}

	return resp
}
