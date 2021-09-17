package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/kampus-merdeka-b/lib/database"
)

func GetArticleController(echoContext echo.Context) error {

	title := echoContext.QueryParam("title")

	articles, err := database.GetArticles(title)
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

	result, err := database.StoreArticle(articleReq.toModel())
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
