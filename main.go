package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

// type person struct {
// 	firstName string
// 	lastName  string
// 	contact contactInfo
// }

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName: "Party",
	// 	contact: contactInfo{
	// 		email: "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }

	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	// jim.updateName("jimmy")
	jim.print()
	// fmt.Printf("%+v", jim)
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}