package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main()  {
	fmt.Println("crud api")

	// get request
	res,err:= http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("error: ",err)
		return
	}

	defer res.Body.Close()

	data,err:= ioutil.ReadAll(res.Body)
	if err!=nil{
        fmt.Println("error: ",err)
        return
    }
	fmt.Println(string(data))
}