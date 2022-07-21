package database

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MaxRetries uint8 = 5
	RetryTime  uint8 = 5
	TimeOut    uint  = 20
)

var (
	ErrInvalidID = errors.New("invalid id")
)

func GetConnection(dsn string) (interface{}, error) {
	var count uint8 = 0
retry:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if MaxRetries == count {
			return db, err
		}
		count++
		fmt.Println("Trying to connect to the database.Retry Count", count)
		time.Sleep(time.Second * time.Duration(RetryTime))
		goto retry
	}
	return db, err
}
