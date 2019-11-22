package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(stringVal(os.Args))
}

func stringVal(v []string) string {
	r := "Hello, world!"
	if len(v) > 1 {
		r = fmt.Sprintf("Hello, %s\n", v)
	}

	return r
}
