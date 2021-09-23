package test

import (
	"github.com/labstack/echo/v4"
	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/models"
)

func InitEcho() *echo.Echo {
	config.InitDB("testing")
	config.MigrateDB()
	e := echo.New()

	return e
}

func TearDown() {
	config.DB.Migrator().DropTable(&models.Article{})
}
