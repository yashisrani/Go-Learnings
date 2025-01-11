package main

import "fmt"

func main()  {
	// if statement
	a:= 10
	if a>5{
		println("a is greater then 10")
	}

	// if-else statements
	b:= 20
	if b>5{
		fmt.Println("b is greater then 5")
	}else{
		fmt.Println("b is not greater then 5")
	}

	// if-else if-else statements
	age:=21
	if age>18{
		fmt.Println("adult")
	}else if age<18{
		fmt.Println("teenager")
	}else{
		fmt.Println("child")
	}

	y:=50
	if y%2 == 0{
        fmt.Println("y is even")
    }else{
        fmt.Println("y is odd")
    }

	if y>5 && (a>5 || b<10){
		fmt.Println("hey how are you ??")
	}else{
		fmt.Println("my name is yash israni")
	}

}