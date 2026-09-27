// go run ./web-service-gin/
// GET
// curl http://localhost:8080/albums
// POST (powershell)
//
//$body = @{
//id = "4"
//title = "The Modern Sound of Betty Carter"
//artist = "Betty Carter"
//price = 49.99
//} | ConvertTo-Json
//
//Invoke-WebRequest `
//    -Uri "http://localhost:8080/albums" `
//-Method POST `
//    -ContentType "application/json" `
//-Body $body

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
