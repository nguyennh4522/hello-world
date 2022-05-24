package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world!")
	})
	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Error:", err)
	}
}
