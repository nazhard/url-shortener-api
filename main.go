package main

import (
	"fmt"
	"github.com/nazhard/url-shortener-api/handler"
	"github.com/nazhard/url-shortener-api/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
		  "message": "Welcome to the URL Shortener API",
		  "create_url": "/shorten",
		})
	})

	r.POST("/shorten", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}