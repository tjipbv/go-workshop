package main

import (
	"fmt"
)

func main() {
	stringVal := "Hello, world!"
	intVal := 192
	boolVal := true

	//fmt.Println(stringVal + intVal + boolVal)

	fmt.Println(stringVal, intVal, boolVal)
	fmt.Printf("%s %d %t\n", stringVal, intVal, boolVal)
}
