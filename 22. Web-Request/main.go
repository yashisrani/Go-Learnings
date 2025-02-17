package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func main()  {
	res,err:= http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("error while getting response", err)
		return
	}
	defer res.Body.Close()

	fmt.Println(res)
	fmt.Printf("data-type of response: %T\n",res)    // *http.Response

	// read the response
	data,err:= ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("error while reading response body", err)
        return
	}
	fmt.Println(string(data))

}