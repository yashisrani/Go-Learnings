package main

import "fmt"


func simplefunction()  {
	fmt.Println("Simple Function")
}

func add(x int , y int) (int) {				// (x (input type), y(input type)) (output type){}
	return x+y
}

func multiple(a, b int) (result int)  {
	result = a*b
	return
}

func main()  {
	simplefunction()
	
	fmt.Println(add(10,20))

	data:= multiple(3,4)
	fmt.Println("multiplication is", data)
}