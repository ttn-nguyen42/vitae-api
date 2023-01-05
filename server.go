package main

import (
	"Vitae/config"
	"Vitae/config/database"
	"Vitae/routes"
	"Vitae/tools/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	server := gin.New()
	database.Connect()
	defer database.Close()
	routes.Create(server)
	port := os.Getenv(config.EnvPort)
	if port == "" {
		port = "8080"
	}
	err := server.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		logging.Fatal(err.Error())
	}
}
