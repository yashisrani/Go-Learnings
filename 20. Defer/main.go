package main

import "fmt"

func main()  {
	// defer keyword is used to run/print code in last ...
	fmt.Println("hi")
    defer fmt.Println("my name is")
	fmt.Println("yash israni")
}