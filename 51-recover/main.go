package main

import (
	"fmt"
)

// recover recovers from the panic so that the cascading effect of panic can be eliminated
// recover must be called using defer. because defer is a must execution block
// recover cannot recover from a fatal
// defer cannot be called when using Fatal from log or any similar pacakges
func main() {
	fmt.Println("Start main")
	defer fmt.Println("End main")
	func() {
		fullName("Jiten", "")
		fmt.Println("Finished calling fullName from main")
	}()
	fullName("Rahim", "Shaik")
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("----->>>", r)
	}
}
func fullName(firstName, lastName string) {
	//defer recoverFullName()
	fmt.Println("Calling fullName func")
	defer fmt.Println("fullName func ends")
	if firstName == "" {
		panic("User Defined Panic:firstName cannot be empty")
		//log.Fatalln("User Defined Fatal:firstName cannot be empty")
	}
	if lastName == "" {
		panic("User Defined Panic:lastName cannot be empty")
		//log.Fatalln("User Defined Fatal:firstName cannot be empty")
	}
	fmt.Println(firstName, lastName)
}
