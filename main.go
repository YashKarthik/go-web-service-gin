package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`      // specifies the field name when struct is serialized into json.
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
    var albumID string = c.Param("id")

    for _, v := range albums {
        if v.ID == albumID {
            c.IndentedJSON(http.StatusOK, v)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }
    albums = append(albums, newAlbum)
    fmt.Println(albums)
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.Run("localhost:8080")
}
