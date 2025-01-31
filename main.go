package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"os"
)


func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
            "email":          "shobayosamuel62@gmail.com",
            "current_datetime": time.Now().UTC().Format(time.RFC3339),
            "github_url":     "https://github.com/Shobayosamuel/hng-task-0",
        })
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}