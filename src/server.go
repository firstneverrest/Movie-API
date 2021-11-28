package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func getMovieLists(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "MovieList",
	})
}

func addNewMovie(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(200, gin.H{"user": u})
}

func updateMovie(c *gin.Context) {
	var user = c.Param("user")
	c.JSON(200, gin.H{"msg": user})
}

func deleteMovie(c *gin.Context) {
	var user = c.Query("user")
	c.JSON(200, gin.H{"msg": user})
}

func main() {
	server := gin.Default()

	server.GET("/", getMovieLists)
	server.POST("/add", addNewMovie)
	server.PUT("/update/:user", updateMovie)
	server.DELETE("/delete", deleteMovie)

	server.Run(":8080")
}