// Package is altijd het eerste statement in Go code files.
// De main package is reserved voor entry level files.
// Dit is de package die opgestart kan worden dmv go run of gebuild via go build.
package main

// Import statements importeren te gebruiken packages. Als je een niet gebruikt package in je import list hebt staan krijg je een imported and not used error
import (
	// fmt is een formatted io function package. String formatting en outputting is gebaseerd op de printf en scanf uit C.
	"fmt"
)

// De main functie is de standaard entrypoint. De application loop start hier.
func main() {
	// fmt.Println print een string (in ons geval hello world) met een OS specifieke newline aan het einde
	fmt.Println("Hello, world!")
}
