package main

import "fmt"

func main()  {
	fmt.Println("Array in Go")

	var name[5] string

	name[0]= "yash"
	name[1] = "mahesh"
	name[2]= "manish"
	name[3]= "manoj"
	name[4]= "anil"

	fmt.Println(name)

	var number = [5]int {1,2,3,4,5}
	fmt.Println(number)
	fmt.Println(len(number))		// to print length of array

	fmt.Println(number[2])
	fmt.Println(name[2])


	var price[5] int
	price[0]= 10
	price[1] = 20
	fmt.Printf("price: %q\n", price)		// %q (quoted string)
	fmt.Printf("name: %q\n", name)
}