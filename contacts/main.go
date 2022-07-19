package main

import (
	"contacts/handlers"
	"flag"

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

	router.GET("/ping", handlers.Ping())
	router.GET("/health", handlers.Health())
	router.GET("/greet", handlers.Greet("Hello World!"))

	//c1 := &models.Contact{ID: 1, Name: "Jiten", Email: "Jitenp@outlook.Com", MoreInfo: "Demo purpose", Address: "Yeshvantpur , Bangalore", ContactNo: "9618558500", Status: "created", LastModified: fmt.Sprint(time.Now().Unix())}

	router.POST("/sample", handlers.SamplePost())
	router.Run(":" + PORT)
}
