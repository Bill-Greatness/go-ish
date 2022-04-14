package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// create an Album Struct
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// create a slice with Album information

var albums = []Album{
	{
		ID:     "1",
		Title:  "Raising Kids",
		Artist: "Samuel Smith",
		Price:  43.2,
	},
	{
		ID:     "2",
		Title:  "Beyond Intelligence",
		Artist: "James Pikes",
		Price:  13.2,
	},
	{
		ID:     "3",
		Title:  "Being You",
		Artist: "Moiraine Rose",
		Price:  0.90,
	},
	{
		ID:     "4",
		Title:  "Religion Poisons Everything",
		Artist: "Christopher Hitchens",
		Price:  50.5,
	},
}

func main() {
	// create a router
	router := gin.Default()

	// using a get request with the router
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)

	// activate router for usage.
	router.Run("localhost:8000")
}

// creating a function to be used in an endpoint
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	// get id from parameter
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// ID not found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
