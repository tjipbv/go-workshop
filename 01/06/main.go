package main

import (
	"fmt"
	"os"
)

func main() {
	val, err := stringVal(os.Args)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println(val)
}

func stringVal(v []string) (string, error) {
	if len(v) < 2 {
		return "", fmt.Errorf("Je hebt geen naam opgegeven")
	}

	return fmt.Sprintf("Hello, %s!", v[1]), nil
}