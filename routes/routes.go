package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/controllers"
)

func NewRoutes() *echo.Echo {
	echoInit := echo.New()
	echoInit.GET("/token", controllers.GetUserToken)
	echoInit.GET("/withToken", controllers.HelloWorld, middleware.JWT([]byte(config.JWT_SECRET)))

	echoInit.GET("/articles", controllers.GetArticleController)
	echoInit.POST("/article", controllers.SaveArticleController)

	return echoInit
}
