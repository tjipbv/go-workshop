package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Name   string `json:"name"`             // struct tag voor json, geliefd en gehaat
	Author string `json:"aUth0r,omitempty"` // json struct tag heeft ook een optie om de variabele niet mee te nemen als hij leeg is
	Year   int    // Geen struct tag format default naar de variabele naam
	id     uint   // Private var wordt niet geoutput maar kan dus wel voor interne doeleinden gebruikt worden.
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		books := []Book{ // Hier maken we een slice van Book
			Book{
				Name:   "1984",
				Author: "George Orwell",
				Year:   1949,
				id:     11,
			},
			Book{ // Author slaan we hier over. In dit geval zal het als leeg worden beschouwd.
				Name:   "Moby Dick", 
				Year:   1851,
				id:     12,
			},
		}

		jsonInBytes, err := json.Marshal(books)
		if err != nil {
			w.Write([]byte("ok"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonInBytes)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
