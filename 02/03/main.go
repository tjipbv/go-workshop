package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string `json:"name"`             // struct tag voor json, geliefd en gehaat
	Author string `json:"aUth0r,omitempty"` // json struct tag heeft ook een optie om de variabele niet mee te nemen als hij leeg is
	Year   int    // Geen struct tag format default naar de variabele naam
	id     uint   // Private var wordt niet geoutput maar kan dus wel voor interne doeleinden gebruikt worden.
}

func main() {
	// Dit kan input zijn uit post body, maar omdat niet iedereen dezelfde tools heeft om met http te praten (bijvoorbeeld curl) even mocken:
	jsonData := []byte(`{"name":"1984","aUth0r":"George Orwell","Year":1949}`)

	book := Book{}
	if err := json.Unmarshal(jsonData, &book); err != nil { // 1 regel bij 1 variabele output. En we passen hier book by reference.
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", book)
}
