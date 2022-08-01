package main

// defer deferes the execution to the end of its caller
// all defers in a stack frame are stacked; That means the first defer is called at the last
func main() {
	defer println("6-Ending Main")
	println("1-Calling Main")
	//yes := new(bool)
	defer func() {
		//*yes = true
		defer println("5-End of anonymous function")
		println("4-Calling a anonymous function")
	}()
	greet()
}

func greet() {
	println("2-Greet is called")
	println("3-Hello World! from greet func")
}
