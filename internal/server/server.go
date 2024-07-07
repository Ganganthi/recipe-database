package server

import (
	"github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
	addLoginRoutes(engine)
}
