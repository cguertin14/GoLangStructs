package main

import (
	"fmt"
)

type ContactInfo struct {
	address string
	email   string
	zipcode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	charles := Person{
		firstName: "Charles",
		lastName:  "Guertin",
		ContactInfo: ContactInfo{
			zipcode: 234234234,
			address: "sdfjlasfj 342 kldfsj",
			email:   "charles@live.ca",
		},
	}

	charles.updateName("Damien")
	charles.print()
}
