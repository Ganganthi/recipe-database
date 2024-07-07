package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"main/internal/database"
	dbModels "main/internal/database/models"
)

func addLoginRoutes(engine *gin.Engine) {
	engine.POST("/signup", createUser)
}

func createUser(ctx *gin.Context) {
	var newUser UserCreate
	if err := ctx.BindJSON(&newUser); err != nil {
		log.Println(err)
		return
	}
	database.DbConn.Create(&dbModels.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Email:     newUser.Email,
	})
	ctx.JSON(http.StatusCreated, gin.H{})
}
