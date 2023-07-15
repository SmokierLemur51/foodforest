package main

import (
	"net/http"

	"github.com/SmokierLemur51/foodforest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", routesindexHandler)

	r.Run(":5000")
}