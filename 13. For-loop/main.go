package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world!")
	}

	// for range statement
	for x := range 11 {
		fmt.Println(x)
	}

	// infinite loop with break statement
	counter:= 0
	for{
		fmt.Println("yash israni")
		counter++
		if counter == 5{
			break
		}
	}


	// for range on slice 
	number:= []int{1,2,3,4,5,6,7,8,9,10,11}
	for index,value := range number {
		fmt.Println(index)
		fmt.Println(value)
	}

	// for range on array
	data:= "hello world!"
	for index, char := range data {
		fmt.Println("Index of data : %d, value is : %c\n", index, char)
	}
}