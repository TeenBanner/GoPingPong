package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Año    int    `json:"año"`
}

var albums = []album{
	{ID: "1", Title: "Famila", Artist: "camila Cablelo", Año: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Año: 2011},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Año: 2022},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado o no existe"})
}

func putAlbumByID(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado o no existe"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("albums/:id", putAlbumByID)
	router.DELETE("albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}
