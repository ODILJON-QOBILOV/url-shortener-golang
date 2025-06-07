package controllers

import (
	"crypto/rand"
	"math/big"

	"github.com/anything/smth/1/config"
	"github.com/anything/smth/1/models"
	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateShortCode(length int) string {
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		code[i] = charset[num.Int64()]
	}
	return string(code)
}

// GetShortedUrl godoc
// @Summary Generate or retrieve a shortened URL
// @Description Accepts a long URL and returns a shortened one. If it exists, returns the existing short URL.
// @Tags urls
// @Accept json
// @Produce json
// @Param request body models.GetUrl true "Original URL"
// @Success 200 {object} map[string]interface{}
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /short-url [post]
func GetShortedUrl(c *gin.Context) {
	var request models.GetUrl
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var db_Url models.Url
	if err := config.DB.Where("url = ?", request.Url).First(&db_Url).Error; err == nil {
		c.JSON(200, gin.H{
			"id":           db_Url.Id,
			"original_url": db_Url.Url,
			"shorted_url":  "https://url-shortener-golang-production.up.railway.app/" + db_Url.ShortedUrl,
		})
		return
	}

	new_Url := models.Url{
		Url:        request.Url,
		ShortedUrl: generateShortCode(5),
	}
	if err := config.DB.Create(&new_Url).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"id":           new_Url.Id,
		"original_url": new_Url.Url,
		"shorted_url":  "https://url-shortener-golang-production.up.railway.app/" + new_Url.ShortedUrl,
	})
}

// GetDataWithShortedUrl godoc
// @Summary Redirect to the original URL
// @Description Redirects the user to the original long URL using the short code.
// @Tags urls
// @Produce plain
// @Param id path string true "Shortened URL ID"
// @Success 301 {string} string "redirect"
// @Failure 404 {object} map[string]string
// @Router /{id} [get]
func GetDataWithShortedUrl(c *gin.Context) {
	id := c.Param("id")
	var url models.Url

	if err := config.DB.Where("shorted_url = ?", id).First(&url).Error; err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(301, url.Url)
}