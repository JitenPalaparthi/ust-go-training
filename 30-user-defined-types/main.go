package main

import "fmt"

func main() {

	emp1 := &Employee{Number: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: Address{Line1: "Flat No 101", Street: "Chinna Swamy Street", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560096"}}
	fmt.Println(emp1)
	fmt.Println("Employee Country", emp1.Address.Country)
	// Promoted fields ?
	student1 := &Student{}
	student1.Number = 101
	student1.Name = "Jiten"
	student1.Email = "JitenP@Outlook.Com"
	student1.Line1 = "Float no 205"
	student1.Street = "Chinnaswamy Street"
	student1.State = "Karnataka"
	student1.Country = "India"
	student1.PinCode = "560096"
	student1.City = "Bangalore"
	fmt.Println(student1)

	// creating object for anonymous fileds
	// This is a good fit if a simple composite type to be created
	p1 := Person{int: 101, string: "Jiten", bool: false}
	fmt.Println(p1)

	p2 := Person{int: 101, string: "Jiten", bool: false}
	fmt.Println(p2)

	if p1 == p2 {
		fmt.Println("Yes p1 and p2 are equal")
	} else {
		fmt.Println("Yes p1 and p2 are not equal")
	}

	emp2 := &Employee{Number: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: Address{Line1: "Flat No 101", Street: "Chinna Swamy Street", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560096"}}
	emp3 := &Employee{Number: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: Address{Line1: "Flat No 101", Street: "Chinna Swamy Street", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560096"}}

	if *emp2 == *emp3 {
		fmt.Println("emp2 and emp3 are equal becase all the fields are equal")
	}

	emp2.Number = 102 // change one of the fields of emp2 then emp2 and emp3 are not equal

	if *emp2 == *emp3 {
		fmt.Println("emp2 and emp3 are equal becase all the fields are equal")
	} else {
		fmt.Println("emp2 and emp3 are not equal becase all the fields are not equal")
	}

	emp4 := Employee{Number: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: Address{Line1: "Flat No 101", Street: "Chinna Swamy Street", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560096"}}
	emp5 := Employee{Number: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Address: Address{Line1: "Flat No 101", Street: "Chinna Swamy Street", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560096"}}

	if emp4 == emp5 {
		fmt.Println("emp4 and emp5 are equal becase all the fields are equal")
	}

	emp4.Number = 102 // change one of the fields of emp2 then emp2 and emp3 are not equal

	if emp4 == emp5 {
		fmt.Println("emp4 and emp5 are equal becase all the fields are equal")
	} else {
		fmt.Println("emp4 and emp5 are not equal becase all the fields are not equal")
	}

	// empty struct
	var e struct{}
	fmt.Println(e)

}

//Note: There is no inheritance in Go
type Employee struct {
	Number      int
	Name, Email string
	Address     Address // Composition
}

type Address struct {
	Line1, Street, City, State, Country, PinCode string
}

type Student struct {
	Number      int
	Name, Email string
	Address     // This is called promoted field
}

// Anonymous fields
type Person struct {
	int
	bool
	string
	//string cannot create two strings
}
