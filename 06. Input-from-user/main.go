package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Println("enter your name:")
	fmt.Scan(&name)

	fmt.Println("name is", name)

	// limitations of fmt.Scan method: it reads until the first whitespace character, so if you want to read a whole line,     then you need to use "bufio" package's NewScanner or ReadString functions for more complex scenarios.

	reader := bufio.NewReader(os.Stdin)
    name, _ = reader.ReadString('\n')
    fmt.Println("hello, ", name)
}
