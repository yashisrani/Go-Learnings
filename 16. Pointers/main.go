package main

import "fmt"

func main()  {
	// Pointer is a variable , which stores the memory address of another variable.
	var num int 
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println(ptr)  // pointer(variable value address)
	fmt.Println(*ptr) // data of variable

	// (2 method)

	name:= "yash"

    pointer:= &name

	fmt.Println(pointer)  // address of variable data
	fmt.Println(*pointer) // data of variable


	var pointer1 *int   // default pointer == nil
	if pointer1== nil{
		fmt.Println("pointer1 is nil")
	}
}