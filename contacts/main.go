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
	flag.StringVar(&DSN, "dbconn", "host=localhost user=postgres password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", `--dbconn= host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`)
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
	router.POST("/sample", handlers.SamplePost())
	cdb := &database.ContactDB{DB: db}
	chandler := &handlers.ContactHandler{IContact: cdb}

	// cfdb := &filedb.ContactFileDB{File: "contacts.fdb"}
	// chandler := &handlers.ContactHandler{IContact: cfdb}

	v1 := router.Group("/v1/contact")
	{
		v1.POST("/create", chandler.CreateContact())
		v1.GET("/get/:id", chandler.GetContactByID())
		v1.PUT("/update/:id", chandler.UpdateContactByID())
		v1.DELETE("/delete/:id", chandler.DeleteContactByID())
	}

	router.Run(":" + PORT)
}
