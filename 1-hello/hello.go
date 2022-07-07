package main

import (
	"fmt"
	"time"
)

func main() {
	println("Hello World")
	t1 := time.Now().Unix()
	t2 := time.Now().Add(24 * time.Hour).Unix()
	fmt.Println((t2 - t1) / 60)
}
