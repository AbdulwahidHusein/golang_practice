package controllers

import (
	"net/http"

	"web-service-gin/models"

	"web-service-gin/store"

	"github.com/gin-gonic/gin"
)

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	store.Albums = append(store.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range store.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.Albums)
}
