package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/xiao0232/pdf2img_gin/handler"
)

func initServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"*"},
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.POST("/upload/pdf", handler.UploadHandler)
	r.Static("/pdf", "./pdf")
	r.Run(":" + port)
}

func main() {
	log.Println("starting server...")

	initServer()
}
