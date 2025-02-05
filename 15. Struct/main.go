package main

import "fmt"

// declaration of a struct
type person struct {
	name   string
	age    int64
	gender string
}

type yash struct {
	height int64
	weight int64
	hair   string
	body   string
	age    int
}

type animal struct {
	name       string
	animaltype string
	food       string
}

type contact struct {
	email string
	phone int
}

type emplopee struct {
	emplopee_details person  // reference to "person" struct
	emplopee_contact contact // reference to "contact" struct
}

func main() {
	// struct(structure) is a composite data type that groups together variables under a single name.
	// each field in a struct can have different data type, and struct are used to create more complex data structures.

	// creating instance of person (1 method)
	var yash person
	yash.age = 21
	yash.gender = "male"
	yash.name = "Yash Israni"

	fmt.Println(yash)

	// (2 method)
	animal := animal{
		name:       "Lion",
		animaltype: "Mammal",
		food:       "Meat",
	}

	fmt.Println(animal)

	// new keyword (3 method)
	var person2 = new(person)
	person2.age = 22
	person2.gender = "female"
	person2.name = "krupali"

	fmt.Println(person2)

	emplopee1 := emplopee{
		emplopee_details: person{
			name:   "John Doe",
			age:    30,
			gender: "male",
		},
		emplopee_contact: contact{
			email: "johndoe@example.com",
			phone: 1234567890,
		},
	}
	fmt.Println(emplopee1)

}
