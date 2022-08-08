package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
	// contact   contactInfo
}

func main() {
	// yury := person{firstName: "Yury", lastName: "Alencar"}

	// fmt.Println("Struct:", yury)
	// fmt.Println("First name:", yury.firstName)
	// fmt.Println("Last name:", yury.lastName)

	ben := person{
		firstName: "Ben",
		lastName:  "Junior",
		contactInfo: contactInfo{
			email:   "benjunior@gmail.com",
			zipCode: 88106444,
		},
	}

	ben.updateName("Benny") // StudyAboutPointers
	ben.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
