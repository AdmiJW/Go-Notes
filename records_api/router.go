package records_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// `getAlbums` responds with the list of all albums in JSON.
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// `postAlbums` adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call `BindJSON` to bind the recieved JSON to `newAlbum`.
	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON provided"})
		logger.Println(err)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// `getAlbumByID` locates the album whose ID value matches the `id`
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON((http.StatusNotFound), gin.H{"message": "album not found"})
}