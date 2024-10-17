package main

import(
	"fmt"
)
func main()  {
	age:=30
	name:="yash"
	height:= 5.10


	// println : to print data in different(new) line
	fmt.Println("age:",age, "height:",height, "name:", name)

	// printf(print formatted) is used for formatted printing.
	
	// There are multiple format specifiers
	// 1) %d = Interger
	// 2) %s = string
	// 3) %T = Type of the value
	// 4) %.3f = float values

	fmt.Printf("age is %d\n", age);
	fmt.Printf("name is %s\n", name);
	fmt.Printf("height is %.3f\n", height);
	fmt.Printf("height is %.4f\n", height);
	fmt.Printf("height is %.2f\n", height);
	fmt.Printf("height is %f\n", height);
	fmt.Printf("Type of age is %T\n", age)

	fmt.Printf("name:%s, age:%d, height:%f\n", name, age, height)
}