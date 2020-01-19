package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", testingPost)
		v1.GET("/", createTodo)
		v1.GET("/:id", createTodo)
		v1.PUT("/:id", putTodo)
		v1.DELETE("/:id", createTodo)
	}
	router.Run()
}

func testingPost(c *gin.Context) {
	fmt.Println("Llego un request!!!")

	completed := c.PostForm("name")
	greeting := c.PostForm("greeting")
	helloResponse := Hello{Greeting: greeting, Name: completed}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": helloResponse})
}

func createTodo(c *gin.Context) {

}

func putTodo(c *gin.Context) {
	var helloJSON Hello
	c.BindJSON(&helloJSON)
	log.Println(helloJSON)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": helloJSON})
}
