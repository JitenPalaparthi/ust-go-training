package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func init() {
	flag.StringVar(&PORT, "port", "58080", "--port=8080")
}

func main() {
	flag.Parse()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"message": "pong"})
		response := ResponseStatus{Message: "Pong", Status: "Success"}
		c.JSON(http.StatusOK, response)
	})

	router.GET("/health", getHealth())

	router.GET("/greet", greet("Hello World!"))

	router.Run(":" + PORT)
}

type ResponseStatus = struct {
	Status  string
	Message string
}

func getHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := ResponseStatus{Message: "OK", Status: "Success"}
		c.JSON(http.StatusOK, response)
	}
}

func greet(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, msg)
	}
}
