package main

import "fmt"

// panic panics and stops execution; the function where panic is called and also cascading stop
// all defer calls are executed before panic.So defer calls are executed irrespective of panics
func main() {
	fmt.Println("Start main")
	defer fmt.Println("End main")
	func() {
		fullName("Jiten", "")
		fmt.Println("Finished calling fullName from main")
	}()
}

func fullName(firstName, lastName string) {
	fmt.Println("Calling fullName func")
	defer fmt.Println("fullName func ends")
	if firstName == "" {
		panic("User Defined Panic:firstName cannot be empty")
	}
	if lastName == "" {
		panic("User Defined Panic:lastName cannot be empty")
	}
	fmt.Println(firstName, lastName)
}
