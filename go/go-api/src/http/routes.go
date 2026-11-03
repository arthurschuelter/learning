package http

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/", GetHealth)
	g.GET("/links", GetLinks)
}
