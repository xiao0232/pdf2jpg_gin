package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/xiao0232/pdf2img_gin/docs"
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

	v1 := r.Group("/api")
	{
		v1.POST("/upload/pdf", handler.UploadHandler)
		// open pdf directory as api endpoint
		v1.Static("/pdf", "./pdf")
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port)
}

// @title pdf2img
// @version 1.0
// @description convert pdf to img(jpg)

// @host localhost:8000
// @BasePath /api

// @license.name Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	log.Println("starting server...")

	initServer()
}
