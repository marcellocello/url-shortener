package handler

import (
	"time"
	"net/http"
	"url-shortener/internal/storage"
	"url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL		string `json:"url" binding:"required"`
	CustomCode	string `json:"custom_code"`
}

func ShortenURL(db *storage.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ShortenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Format JSON salah atau URL Kosong"})
			return
		}

		shortCode := req.CustomCode
		if shortCode == "" {
			shortCode = utils.GenerateShortCode()
		}

		_, err := db.Conn.Exec("INSERT INTO urls (original_url, short_code) VALUES ($1, $2)", req.URL, shortCode)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Custom Code sudah digunakan, cari yang lain!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"origin_url": req.URL,
			"short_url": "http://localhost:8080/r/" + shortCode,
		})
	}
}

func RedirectURL(db *storage.DB, rdb *storage.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Param("code")
		ctx := c.Request.Context()

		val, err := rdb.Client.Get(ctx, code).Result()
		if err == nil {
			c.Redirect(http.StatusFound, val)
			go db.Conn.Exec("UPDATE urls SET visit_count = visit_count + 1 WHERE short_code = $1", code)
			return
		}

		var originalURL string
		err = db.Conn.QueryRow(
			"UPDATE urls SET visit_count = visit_count + 1 WHERE short_code = $1 RETURNING original_url",
			code,
		).Scan(&originalURL)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL tidak ditemukan"})
			return
		}

		rdb.Client.Set(ctx, code, originalURL, 24*time.Hour)

		c.Redirect(http.StatusFound, originalURL)
	}
}
