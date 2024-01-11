package main

import "fmt"
import "encoding/json"

type Person struct {
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Address string `json:"Address"`
}

func main() {
	myJson := `[
		{
			"Name": "John",
			"Age": 30,
			"Address": "New York"
		},
		{
			"Name": "Jane",
			"Age": 25,
			"Address": "London"
		}
	]`

	// Unmarshal JSON to struct
	var unmarshalled []Person

	// Unmarshal takes slice of bytes (strings are slices of bytes) and a pointer to the struct (interface{})
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(unmarshalled)

	// write json from struct
	var mySlice []Person

	var m1 Person
	m1.Name = "Ryan"
	m1.Age = 24
	m1.Address = "Daly City"

	mySlice = append(mySlice, m1)

	// MarshalIndent takes value, prefix, and indent
	newJson, err := json.MarshalIndent((mySlice), "", "    ")

	fmt.Println(string(newJson))
}
