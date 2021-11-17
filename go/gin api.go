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
	// Initialize Gin Router
	router := gin.Default()

	// Create endpoints for the API.
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	// Attach the router to an http server, and start it
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	// deserialize the JSON into the album struct.
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// add the album to the slice of albums.
	albums = append(albums, newAlbum)

	// serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// curl -X POST -H "Content-Type: application/json" -d '{"id":"4","title":"The Beatles","artist":"The Beatles","price":10.99}' http://localhost:8080/albums	// Create a new album

// getAlbumById responds with the album with the specified ID as JSON.
func getAlbumById(c *gin.Context) {
	// get the id from the URL.
	id := c.Param("id")

	// find the album with the matching id.
	for _, album := range albums {
		if album.ID == id {
			// serialize the struct into JSON and add it to the response.
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// if no album is found, return a 404.
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "album not found"})
}
