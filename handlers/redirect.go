package handlers

import (
  "github.com/nazhard/url-shortener-api/store"
  "github.com/gin-gonic/gin"
)

func HandleShortUrlRedirect(c *gin.Context) {
  shortUrl := c.Param("shortUrl")
  initialUrl := store.RetrieveInitialUrl(shortUrl)
  c.Redirect(302, initialUrl)
}