package main

import (
	"fmt"
	"strings"
)

func main()  {

	// strings.Split() function splits a string into a slice of strings by a given separator.
	data:= "apple,bannana,orange"
	parts:= strings.Split(data, ",")    // (variable name, "sparator")
	fmt.Println(parts)

	// strings.Contains() function checks if a string contains another string.
	str:=  "one two three four five two five five"
	count:= strings.Count(str, "five")
	fmt.Println(count)

	// strings.TrimSpace() function removes all leading and trailing white spaces from a string.
	yash:= "     hi yash, how are you ??     "
	trimmed:= strings.TrimSpace(yash)
	fmt.Println(trimmed)

	// strings.Join() function concatenates all strings in a slice into a single string, separated by the specified separator.
	str1:= "yash"
	str2:= "israni"
	join:= strings.Join([]string{str1,str2}," ")
	fmt.Println(join)
}