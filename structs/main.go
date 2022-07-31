package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// yury := person{firstName: "Yury", lastName: "Alencar"}

	// fmt.Println("Struct:", yury)
	// fmt.Println("First name:", yury.firstName)
	// fmt.Println("Last name:", yury.lastName)

	ben := person{
		firstName: "Ben",
		lastName:  "Junior",
		contact: contactInfo{
			email:   "benjunior@gmail.com",
			zipCode: 88106444,
		},
	}

	fmt.Printf("%+v", ben)
}
