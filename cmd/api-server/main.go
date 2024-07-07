package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"main/internal/database"
	"main/internal/server"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	engine := gin.Default()

	server.AddRoutes(engine)

	engine.Run()
}
