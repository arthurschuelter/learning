package http

import (
	"net/http" // Study
	// "context" 	// Study
	// "fmt"      // Study
	// "time"			// Study

	"github.com/gin-gonic/gin"
)

var test int

func Configure() {
	test = 1
}

func GetHealth(ctx *gin.Context) {
	response := gin.H{
		"message": "The api is healthy",
	}
	ctx.JSON(http.StatusOK, response)
}

func GetLinks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, test)
}
