package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kinneko-de/restaurant-document-design-gateway/build"
	"github.com/kinneko-de/restaurant-document-design-gateway/internal/app/document"
)

func main() {
	log.Println("Version " + build.Version)
	router := setupRouter()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/document/preview", document.GeneratePreview)
	return router
}
