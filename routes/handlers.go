package routes

import (
	"github.com/gin-gonic/gin"
)



func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Food Forest"
	})
}