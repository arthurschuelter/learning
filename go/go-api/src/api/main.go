package main

import (
	"go-api/src/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	http.Configure()
	http.SetRoutes(g)
	if err := g.Run(":3000"); err != nil {
		panic(err)
	}
}
