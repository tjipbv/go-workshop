package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// HandleFunc is een http handler die we verbinden aan het pad / (en automatisch alles wat er achter zit)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Hier halen we via het request object de gegeven get parameter met key name op
		v, err := stringVal(r.URL.Query().Get("name"))
		if err != nil {
			// Als de stringVal functie een foutmelding geeft sturen we een status code 418 terug met als body de error message
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte(err.Error()))
			return
		}
		// En als er geen error message is de uitkomst van stringval
		w.Write([]byte(v))
	})

	// Hier zeggen we: start de server op poort 8080 en als het foutgaat wil ik graag dat je de applicatie stopt en de error naar stderr doorzet
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// v []string word nu v string. we gaan een query parameter doorgeven, niet een arglist
func stringVal(v string) (string, error) {
	if v == "" {
		return "", fmt.Errorf("Je hebt geen naam opgegeven")
	}

	return fmt.Sprintf("Hello, %s!", v), nil
}

