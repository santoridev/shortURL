package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)

func urlShorter(c *gin.Context) {
	var request struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
	}

	shortKey := RandString(5)
	urlMap[shortKey] = request.URL

	c.JSON(http.StatusOK, gin.H{"short_url": "/" + shortKey})
}

func urlRedirect(c *gin.Context) {
	key := c.Param("shorturl")
	longurl, exists := urlMap[key]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, longurl)

}
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {

	r := gin.Default()

	r.POST("/shorten", urlShorter)
	r.GET("/:shorturl", urlRedirect)

	r.Run(":8080")
}
