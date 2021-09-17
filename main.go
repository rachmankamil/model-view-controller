package main

import (
	"github.com/rachmankamil/kampus-merdeka-b/config"
	"github.com/rachmankamil/kampus-merdeka-b/routes"
)

func main() {
	config.InitDB()
	echoApp := routes.NewRoutes()
	echoApp.Start(":8080")
}
