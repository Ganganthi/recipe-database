package login

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main/internal/database"
	dbModels "main/internal/database/models"
)

func AddRoutes(engine *gin.Engine) {
	engine.POST("/signup", createUser)
}

func createUser(ctx *gin.Context) {
	newUser := dbModels.User{
		FirstName: "abc",
		LastName:  "def",
		Username:  "user",
		Password:  "pass",
		Email:     "email@test.com",
	}
	database.DbConn.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{})
}
