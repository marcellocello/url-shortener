package routes

import (
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth_gin"
)

func SetupRouter(db *storage.DB, rdb *storage.RedisClient) *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	limiter := tollbooth.NewLimiter(1, nil)

	r.GET("/", func(c *gin.Context) {
    		rows, err := db.Conn.Query("SELECT short_code, visit_count FROM urls ORDER BY visit_count DESC LIMIT 5")
    		if err != nil {
        		c.HTML(http.StatusOK, "index.html", gin.H{
            			"Links": []interface{}{},
            			"Error": "Gagal mengambil data dari database",
        		})
        		return
    		}

   		 defer rows.Close()

    		type LinkInfo struct {
        		ShortCode  string
        		VisitCount int
    		}
    		var links []LinkInfo

    		for rows.Next() {
        		var l LinkInfo
        		if err := rows.Scan(&l.ShortCode, &l.VisitCount); err != nil {
            			continue
        		}
        		links = append(links, l)
    		}

  	  	c.HTML(http.StatusOK, "index.html", gin.H{
        		"Links": links,
    		})
	})

	r.POST("/shorten", tollbooth_gin.LimitHandler(limiter), handler.ShortenURL(db))
	r.GET("/r/:code", handler.RedirectURL(db, rdb))

	return r
}
