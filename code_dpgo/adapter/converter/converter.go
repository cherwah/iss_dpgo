package converter

import (
	"encoding/json"
	"strings"
)

type Convertible interface {
	Convert(dict string) map[string]string
}

type Csv struct {
}

type Json struct {
}

func (c *Csv) Convert(csv string) map[string]string {
	// create an empty dictionary
	dict := map[string]string{}

	// split elements separated by ","
	str := strings.Split(csv, ",")
	for _, pair := range str {
		// split pairs separated by ":"
		name_val := strings.Split(pair, ":")
		dict[name_val[0]] = name_val[1]
	}

	return dict
}

func (j *Json) Convert(json_str string) map[string]string {
	// create an empty dictionary
	dict := map[string]string{}

	json.Unmarshal([]byte(json_str), &dict)

	return dict
}
