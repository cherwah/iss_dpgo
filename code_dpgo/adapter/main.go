package main

import (
	"adapter/converter"
	"fmt"
)

/*
Adapting objects that output different formats
to use the same logging function.
*/
func main() {
	// adapting to use log()
	csv := "name:john,age:23,gender:m"

	csv_adapter := converter.Csv{}
	dict1 := csv_adapter.Convert(csv)
  // adapted our data to work with existing interface
	log(dict1)

	// adapting to use log()
	json := `{"name":"john","age":"23","gender":"m"}`

	json_adapter := converter.Json{}
	dict2 := json_adapter.Convert(json)a
  // adapted our data to work with existing interface
	log(dict2)
}

// the log() function only works with dictionaries
func log(dict map[string]string) {
	for key, val := range dict {
		fmt.Printf("%s: %s\n", key, val)
	}

	fmt.Println()
}
