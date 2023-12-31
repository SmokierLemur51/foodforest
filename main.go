package main

import (
	"github.com/SmokierLemur51/munchmania/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", routes.IndexHandler)

	r.Run(":5000")
}