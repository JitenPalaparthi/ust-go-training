package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/greet", greet("Hello World!"))

	fmt.Println("Server started; running on port 50080")
	http.ListenAndServe(":50080", nil)
	// 0.0.0.0
	// what is 0.0.0.0 --> All interfaces in the machine
	// lo --> loopback  that is 127.0.0.1 a.k.a localhost
	// wl --> can be wireless interface connected to the router
	// bluetooth
	// et0
	// veth
	// docker bridge
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")

}

func greet(msg string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
	}
}
