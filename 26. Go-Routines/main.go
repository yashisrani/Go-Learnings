package main

import (
	"fmt"
	"time"
)

func main()  {
	// go routines

	// every process which runs concurrently in golang called go routines
	// go routes = light weight threads
	// creation cost of go routines is very small as compared with threads
	// every program had atleast one go routine which called as main function
	// when main function is terminated , then all go routines will terminated means all go routines works under main function.

	// syntax of go routines
	// func function_name(){
	 //     // code here
     // }
     // go function_name()

    
	go doSomething() // create and start go routine
	doSomething()

	go func ()  {		// anonymous function
		fmt.Println("hi there")
	}()
	time.Sleep(1*time.Second)
}

func doSomething()  {
	
	for i := 0; i <=10; i++ {
		fmt.Println("hello from go routine", i)
        time.Sleep(1*time.Second) // delay for 1 second
	}
	
}

