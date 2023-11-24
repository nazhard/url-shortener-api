package main

import (
  "fmt"
  
  "github.com/nazhard/url-shortener-api/handlers"
  "github.com/nazhard/url-shortener-api/store"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "create_url": "/shorten",
      "message": "Welcome to the URL Shortener API",
    })
  })
  
  r.POST("/shorten", func(c *gin.Context) {
    handlers.CreateShortUrl(c)
  })
  
  r.GET("/:shortUrl", func(c *gin.Context) {
    handlers.HandleShortUrlRedirect(c)
  })
  
  // Note store initialization happens here
  store.InitializeStore()
  err := r.Run(":9808")
  if err != nil {
    panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
  }
}