package records_api

import (
	"log"

	"github.com/gin-gonic/gin"
)

var logger = log.Default()

func Init() {}

func Main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
