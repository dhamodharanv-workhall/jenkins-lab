package main
import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

var app int8 = 20

// algebric data type
type album struct {
  ID string `json:"id"` // while doing the json serialization this value is considered under the name "id".
  Title string `json:"title"`
}

// slice wiht album data type
var albums = []album{
  {ID: "1", Title: "lunch"},
}

//getAlbums and responds
func getAlbums(c *gin.Context)  {
  c.IndentedJSON(http.StatusOK, albums)
}

func main() {
  fmt.Println("Simple Server.")
  router := gin.Default()
  router.GET("/albums", getAlbums)
  router.Run("0.0.0.0:8000")
}

