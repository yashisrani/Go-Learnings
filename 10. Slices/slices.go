package main

import "fmt"

func main()  {
	fmt.Println("Slices in Go")

	// Slice is flexible and dynamic data structure that provides a more powerful alternative to arrays

	// declaration of slice
	name:= []string{"yash","manish","mahesh","manoj","vinod"}

	// to add more data in slice
	name = append(name, "vikas","vishal", "raj","raju","mighty raju")

	// access and print slice elements
	fmt.Println(name)

	// length of slice
	fmt.Println(len(name))

	// capacity of slice
	fmt.Println(cap(name))

	// type of slice
	fmt.Printf("datatype: %T\n", name)

	// when we want to declare a slice without providing any initial value , then we need to use the "make function" to create slice with specified length and capacity

	var numbers= make([]int, 12, 30)   // make([]type , length, capacity)

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}