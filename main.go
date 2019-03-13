package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
}

func main() {
	// Three ways to assign values to a struct in Go

	// 1.
	// alex := person{"Alex", "Anderson"}

	// 2.
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson"}

	// 3.
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")

	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(p).firstName = newFirstName
}
