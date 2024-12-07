package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album representa dados sobre um álbum de música.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice para armazenar dados dos álbuns.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responde com a lista de todos os álbuns em formato JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adiciona um álbum a partir do JSON recebido no corpo da requisição.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Chama BindJSON para associar o JSON recebido ao novo álbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Adiciona o novo álbum à lista de álbuns.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID localiza o álbum cujo ID corresponde ao parâmetro id
// enviado pelo cliente e retorna esse álbum como resposta.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop pelos álbuns, procurando um álbum cujo ID corresponda ao parâmetro.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}