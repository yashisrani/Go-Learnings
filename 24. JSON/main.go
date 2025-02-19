package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	// JSON is a lightweight data-interchange format that is easy for humans to read and write and easy for machines to parse and generate.

	// example to demonstrate "how to define and use json in golang ? "
	// 1) create a struct
	// 2) Marshalling (Encoding)
	// 3) Unmarshalling (Decoding)

	person:= Person{Name:"yash", Age: 21, Gender: "male"}
	fmt.Println(person)

	//  Encoding (Marshalling) (to convert person into json format) 
	jsondata,err:= json.Marshal(person)
	if err!=nil{
        fmt.Println("error while encoding json", err)
        return
    }
	fmt.Println(string(jsondata))  // if we want to print the data then we want to convert this data in readble format(string format)

	// decoding (Unmarshalling)
	var decodedperson Person
	error1:= json.Unmarshal(jsondata, &decodedperson)
	if error1!=nil{
        fmt.Println("error while decoding json", error1)
        return
    }
	fmt.Println(decodedperson) // no need to convert this data in readble format(string format)
}
