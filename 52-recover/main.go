package main

// fatals cannot be recovered
func recovers() {
	if r := recover(); r != nil {
		println(r)
	}

}
func main() {
	done := make(chan bool)
	defer recovers()
	done <- true
	<-done
	println("Finished calling main")
}
