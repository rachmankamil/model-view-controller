package main

import (
	"github.com/rachmankamil/kampus-merdeka-b/config"
	ownMid "github.com/rachmankamil/kampus-merdeka-b/middleware"
	"github.com/rachmankamil/kampus-merdeka-b/routes"
)

func main() {
	config.InitDB("")

	echoApp := routes.NewRoutes()
	ownMid.LogMiddlewareInit(echoApp)
	echoApp.Start(":8080")
}
