package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//

type Person struct {
	ID           int
	Name         string
	Email        string
	Status       string
	LastModified string
}

type Student struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "host=localhost user=admin password=admin123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)

	db.AutoMigrate(&Person{})
	db.Create(&Person{Name: "Jiten", Email: "JitenP@Outlook.Com", Status: "created", LastModified: time.Now().GoString()})
	db.AutoMigrate(&Student{})
	db.Create(&Student{Name: "Jiten", Email: "JitenP@Outlook.Com"})
	p1 := &Person{}
	s1 := &Student{}
	db.First(p1, 1)
	db.First(s1, 1)
	fmt.Println(p1)
	fmt.Println(s1)
}
