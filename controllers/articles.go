package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/kampus-merdeka-b/lib/database"
	"github.com/rachmankamil/kampus-merdeka-b/lib/upload"
)

func GetArticleController(echoContext echo.Context) error {

	title := echoContext.QueryParam("title")

	articles, err := database.GetArticles(title, 0, 0)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponseArray(*articles),
	})
}

func SaveArticleController(echoContext echo.Context) error {
	var articleReq RequestArticle
	echoContext.Bind(&articleReq)

	file, err := echoContext.FormFile("file")
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	path, err := upload.UploadLocal(file, "articles")
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	articleModel := articleReq.toModel()
	articleModel.ThumbnailPath = path

	result, err := database.StoreArticle(articleModel)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   newResponse(*result),
	})
}
