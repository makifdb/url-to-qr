package main

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/makifdb/url-to-qr/converter"
)

//go:embed all:frontend/dist
var distfs embed.FS

func StaticsFS() embed.FS {
	return distfs
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()
	app.Use(gin.LoggerWithFormatter(func(ginlog gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s\"\n",
			ginlog.TimeStamp.Format("2006-01-02 15:04:05"),
			ginlog.ClientIP,
			ginlog.Method,
			ginlog.Path,
			ginlog.Request.Proto,
			ginlog.StatusCode,
			ginlog.Latency,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(CORSMiddleware())

	app.GET("/assets/*filepath", func(c *gin.Context) {
		c.FileFromFS(path.Join("/frontend/dist/", c.Request.URL.Path), http.FS(StaticsFS()))
	})

	app.Any("/", func(c *gin.Context) {
		c.FileFromFS("frontend/dist/", http.FS(StaticsFS()))
	})

	app.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("frontend/dist/favicon.ico", http.FS(StaticsFS()))
	})

	app.POST("/api/v1/convert", createQR())

	app.NoRoute(func(c *gin.Context) {
		c.FileFromFS("frontend/dist/", http.FS(StaticsFS()))
	})

	log.Println("Starting server on http://localhost:3000")

	if err := app.Run(net.JoinHostPort("", "3000")); err != nil {
		log.Fatal(err)
	}
}

func createQR() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the url from the request
		var req struct {
			URL     string `json:"url"`
			BGColor string `json:"bg_color"`
			FGColor string `json:"fg_color"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// check if the url is valid
		if req.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
			return
		}

		image, err := converter.CreateQRCode(req.URL, req.BGColor, req.FGColor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// return the image
		c.JSON(http.StatusOK, gin.H{"image": image})
	}
}
