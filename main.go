package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultMesage = "Hello, World!"

func main()  {
	engine := gin.Default()
	
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, defaultMesage)
	})
	
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	
	engine.Run(":8080")
}
