package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddlewareInit(e *echo.Echo) {
	logger := middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
		` + "\n",
	}

	e.Use(middleware.LoggerWithConfig(logger))
}
