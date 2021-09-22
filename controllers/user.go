package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	ownMid "github.com/rachmankamil/kampus-merdeka-b/middleware"
)

func GetUserToken(echoContext echo.Context) error {

	username := echoContext.QueryParam("username")
	password := echoContext.QueryParam("password")

	token := ""
	var err error
	if username == "kamil" && password == "masukaja" {
		token, err = ownMid.CreateToken(1)
	}

	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "error",
			"data":   err,
		})
	} else {
		return echoContext.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   token,
		})
	}
}

func HelloWorld(echoContext echo.Context) error {
	return echoContext.JSON(http.StatusAccepted, "Logged in")
}
