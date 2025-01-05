package main

import "fmt"

func updatenumber(x *int) int {
	*x = 10
	return *x
}
func main()  {
	x:=5
	fmt.Println(x)

	updatenumber(&x)	// pass by reference
	fmt.Println(x)
}