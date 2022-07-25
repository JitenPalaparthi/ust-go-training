package main

import (
	"contacts/database"
	ech "contacts/handlers/echo"
	gn "contacts/handlers/gin"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

var (
	PORT   string
	DSN    string
	BROKER string
	FW     string
)

func init() {
	flag.StringVar(&PORT, "port", "58080", "--port=8080")
	flag.StringVar(&DSN, "dbconn", "host=localhost user=postgres password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", `--dbconn= host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`)
	flag.StringVar(&BROKER, "brokers", "localhost:29092", "--brokers=localhost:29092,localhost:29093")
	flag.StringVar(&FW, "webframework", "gin", "--webframework=echo")

}

func main() {

	flag.Parse()

	db, err := database.GetConnection(DSN)
	fmt.Println(db, err)
	if err != nil {
		panic(err)
	}
	switch FW {
	case "gin", "GIN":
		runGinFramework(PORT, db)
	case "echo", "ECHO":
		runEchoWebFramework(PORT, db)
	case "both", "BOTH":
		if err != nil {
			return
		}
		port, err := strconv.Atoi(PORT)
		if err != nil {
			return
		}
		port = port + 1
		go runGinFramework(PORT, db)
		runEchoWebFramework(strconv.Itoa(port), db)
	default:
		fmt.Println("provided framework has not implemented")
	}
}

func runGinFramework(port string, db interface{}) {
	router := gin.Default()
	router.GET("/ping", gn.Ping())
	router.GET("/health", gn.Health())
	router.GET("/greet", gn.Greet("Hello World!"))
	router.POST("/sample", gn.SamplePost())
	brokers := strings.Split(BROKER, ",")
	cdb := &database.ContactDB{DB: db}
	chandler := &gn.ContactHandler{IContact: cdb, Brokers: brokers}

	// cfdb := &filedb.ContactFileDB{File: "contacts.fdb"}
	// chandler := &gin.ContactHandler{IContact: cfdb}

	v1 := router.Group("/v1/contact")
	{
		v1.POST("/create", chandler.CreateContact())
		v1.GET("/get/:id", chandler.GetContactByID())
		v1.PUT("/update/:id", chandler.UpdateContactByID())
		v1.DELETE("/delete/:id", chandler.DeleteContactByID())
	}
	router.Run(":" + port)
}

func runEchoWebFramework(port string, db interface{}) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	brokers := strings.Split(BROKER, ",")
	cdb := &database.ContactDB{DB: db}
	chandler := &ech.ContactHandler{IContact: cdb, Brokers: brokers}
	contactGroup := e.Group("/v1/contact")
	contactGroup.POST("/create", chandler.CreateContact())
	contactGroup.GET("/get/:id", chandler.GetContactByID())
	contactGroup.PUT("/update/:id", chandler.UpdateContactByID())
	contactGroup.DELETE("/delete/:id", chandler.DeleteContactByID())
	e.Logger.Fatal(e.Start(":" + port))
}
