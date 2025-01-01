package main

import "fmt"

// main function
func main()  {
	
	// variable with specified datatype
	var name string = "yash israni"
	fmt.Printf(name)

	var money int = 10000
	fmt.Println(money)

	var isStudent bool = true
	fmt.Println(isStudent)

	var grade float64 = 85.5
	fmt.Println(grade)


	// variable with inferred datatype
	var name1 = "yash israni"
	fmt.Printf(name1)


	// variable with short declaration
	name2:= "yash israni"
	fmt.Printf(name2)

	// declare and assign two variable at same time
	contact, marks := 193891,99.9
	fmt.Println(contact,marks)

	// const variable
	const pi = 3.14
	fmt.Println(pi)

	// if first letter of name is "Capital", then it's exported(public). it's called "Public Variables"
	var Public = "data is important"

	// if first letter of name is "not Capital", then it's unexported(private). it's called "Private Variables"
	var private = "data is private"

	fmt.Println(Public)
	fmt.Println(private)

	// public function
	func PublicFuction(){
		fmt.Println("public function")
	}

	PublicFunction();
	
	// private function
	func privatefunction(){
		fmt.Println("private function")
	}
}