package main

import (
	"fmt"
	"strconv"
)

func main()  {
	num:=42
	fmt.Println(num)
	fmt.Printf("Type %T\n", num)	// data-type : int

	data:= float64(num)
	fmt.Println(data)
	fmt.Printf("type %T\n", data)    // data-type : float

	number:= 10
	str:= strconv.Itoa(number)
	fmt.Println(str)
	fmt.Printf("data-type is %T\n", str)  // data-type: string

	name:= "yashisrani"
	name_int,_:= strconv.Atoi(name) // Atoi returns to values(int,error) , so we don't want error value .. then we add _ 
	fmt.Println(name_int) // output: 0
	fmt.Printf("data-type is %T\n", name_int)  // data-type : int

	a:="3.14"
	b,_:= strconv.ParseFloat(a,64)
	fmt.Println(b)
	fmt.Printf("data-type %T\n", b)

}
