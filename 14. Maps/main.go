package main

import "fmt"

func main()  {

	// map is a data structure that provides an unordered collection of key-value pairs, where each key must be unique.
	// it is similar to dictionary and hashmaps in other programming languages.

	// declaration of map (student-name <-> student-marks)
	studentgrades:= make(map[string]int)

	// adding key-value pairs
	studentgrades["yash israni"] = 85
	studentgrades["john doe"] = 90
	studentgrades["mahesh"] = 10
	studentgrades["manu"] = 100

	// accessing values using keys
	fmt.Println(studentgrades["yash israni"])

	// modifying values using keys
	studentgrades["manu"] = 20

	// deleting key-value pairs
	delete(studentgrades, "mahesh")

	// checking key is exists or not
	grades, exists:= studentgrades["yash israni"]
	fmt.Println(exists)
	fmt.Println(grades)

	for index,value := range studentgrades {
		fmt.Println(index)
		fmt.Println(value)
	}

	// initializing map with declaration
	rollno:= map[int]string{
		1: "yash israni",
        2: "john doe",
        3: "mahesh",
        4: "manu",
	}

	for index, value := range rollno {
		fmt.Printf("rollno is: %d and name is %s\n", index, value)
	}
}