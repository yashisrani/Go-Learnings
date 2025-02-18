package main

import (
	"fmt"
	"net/url"
)

func main()  {
	// URLs (Uniform Resource Locators)

	myurl:= "http://localhost:8080/admin"
	fmt.Println(myurl)
	fmt.Printf("type: %T\n",myurl)  // type: string


	// parse the URL to get individual components like scheme, host, path etc. (to convert string into url)
	parsedurl,err:= url.Parse(myurl)	
	fmt.Println(parsedurl)
	if err!=nil{
		fmt.Println("error while parsing url", err)
        return
	}
	fmt.Printf("data-type: %T\n",parsedurl)		// data-type: *url.URL

	fmt.Println("scheme: ",parsedurl.Scheme)
	fmt.Println("host: ",parsedurl.Host)
	fmt.Println("path: ",parsedurl.Path)
	fmt.Println("Raw query: ",parsedurl.RawQuery)

	// modify the parsed URL components
	parsedurl.Path = "/newPath"
	parsedurl.RawQuery = "username:yash"

	// construct the url 
	newurl:= parsedurl.String()
	fmt.Println("new url: ",newurl)   //  http://localhost:8080/newPath?username:yash
} 