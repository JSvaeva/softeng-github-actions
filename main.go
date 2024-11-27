package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, Hello())
	})

	router.Run(":3000")
}