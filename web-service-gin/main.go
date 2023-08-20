package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	// Endpoints
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

/*
*

	getAlbums responds with the list of all albums as JSON.
	Note: You could have given this function any name.
	Neither Go nor Gin cares what you call it/requires a specific function name format/
*/
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	// The request body must contain a JSON object with keys that match the album fields.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	// In the real world, you would probably use a database of some kind.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// If no album is found, return an empty album with a 404 status code.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
