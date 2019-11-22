package main

import (
	"fmt"
	"os" // os package voor interface met os
)

func main() {
	stringVal := "Hello, world!"
	if len(os.Args) > 1 {
		stringVal = fmt.Sprintf("Hello, %s\n", os.Args[1])
	}

	fmt.Println(stringVal)
}
