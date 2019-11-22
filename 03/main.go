package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseUrl string = "https://go.jegezicht.nl/draw"

type DrawInput struct {
	X     uint   `json:"x"`
	Y     uint   `json:"y"`
	Color string `json:"color"`
}

func main() {

	drawing := DrawInput{
		X:     50,
		Y:     40,
		Color: "red",
	}

	if err := call(drawing); err != nil {
		fmt.Println(err)
	}
}

func call(drawing DrawInput) error {

	w := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(w).Encode(drawing); err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", baseUrl, w)
	if err != nil {
		return err
	}

	req.Header.Add("IkBen", ``)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
