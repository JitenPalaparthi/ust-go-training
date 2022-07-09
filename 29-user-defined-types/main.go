package main

import "fmt"

func main() {
	var emp1 Employee // declare
	emp1 = Employee{} // instantiate
	fmt.Println(emp1) // Type inference work here. All number types gets 0 and bool
	emp1.Number = 100
	emp1.Name = "JP"
	emp1.IsMarried = false
	emp1.Email = "Jiten.Palaparthi@Gmail.Com"
	emp1.Salary = 10000.45
	emp1.Address = "Dilsukhnagar, Hyderabad"
	fmt.Println(emp1)
	emp2 := Employee{Number: 101, IsMarried: true, Salary: 12000.50, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: "Rajajinagar,Bangalore"}
	fmt.Println(emp2)
	// the way the reference type of struct varialbe is same as normal one
	// create a pointertype using new function
	emp3 := new(Employee)
	fmt.Println(emp3)
	emp3.Number = 102
	emp3.Name = "Jayan"
	emp3.IsMarried = true
	emp3.Email = "Jayan.K@Gmail.Com"
	emp3.Salary = 18000.45
	emp3.Address = "Medical College, Trivandrum"
	fmt.Println(emp3)

	// Shorthand declaration to create a reference type
	emp4 := &Employee{}
	fmt.Println(emp4)
	emp4.Number = 103
	emp4.Name = "Rasool"
	emp4.IsMarried = true
	emp4.Email = "Rasool.Sheik@Gmail.Com"
	emp4.Salary = 18000.45
	emp4.Address = "Gandhi Nagar, Hyderabad"
	fmt.Println(emp4)
}

// To create a user defined type we can use type keyword
// To crete a struct

type Employee struct {
	Number    int
	IsMarried bool
	Salary    float32
	Name      string
	Email     string
	Address   string
}
