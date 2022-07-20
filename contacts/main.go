package main

import (
	"contacts/database"
	"contacts/handlers"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func init() {
	flag.StringVar(&PORT, "port", "58080", "--port=8080")
	flag.StringVar(&DSN, "dbconn", "host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", `--dbconn= host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`)
}

func main() {
	flag.Parse()
	router := gin.Default()

	db, err := database.GetConnection(DSN)
	fmt.Println(db, err)
	if err != nil {
		panic(err)
	}
	router.GET("/ping", handlers.Ping())
	router.GET("/health", handlers.Health())
	router.GET("/greet", handlers.Greet("Hello World!"))

	//c1 := &models.Contact{ID: 1, Name: "Jiten", Email: "Jitenp@outlook.Com", MoreInfo: "Demo purpose", Address: "Yeshvantpur , Bangalore", ContactNo: "9618558500", Status: "created", LastModified: fmt.Sprint(time.Now().Unix())}

	router.POST("/sample", handlers.SamplePost())
	router.Run(":" + PORT)
}
