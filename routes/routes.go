package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/kampus-merdeka-b/controllers"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()

	echoInit.GET("/article", controllers.GetArticleController)
	echoInit.POST("/article", controllers.SaveArticleController)

	return echoInit
}
