package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"main/internal/database"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong4",
		})
	})
	r.Run()
}
