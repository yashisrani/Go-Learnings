package main

import (
	// "bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	file, err:=os.Create("example.txt")
	if err != nil {
		fmt.Println("error while creating ",err)
		return
	}
	defer file.Close()

	content:= "my name is yash israni"
	bytes,err:= io.WriteString(file,content)
	fmt.Println(bytes)

	// read the file contents into buffer

	// create a buffer to read the content
	buffer:= make([]byte,1024)

	// read the file line by line
	for{
		n,err :=file.Read(buffer)	// returns n(variable) and err(error)
		if err == io.EOF{	// EOF (end of file)
			break
		}
		if err !=nil{
			fmt.Println("error while reading ",err)
            return
		}

		// process the read content
		fmt.Println(string(buffer[:n]))  // buffer read 0:n(0 to n)
	}

	// (shortcut method) to read the entire file into a byte slice

	// this method is useful to read entire file at onces ..
	
	con,error:= ioutil.ReadFile("example.txt")
	if error!=nil{
		fmt.Println("error while reading ",error)
		return
	}
	fmt.Println(string(con))

}