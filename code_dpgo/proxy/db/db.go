package db

import (
	"encoding/json"
	"os"
)

type Json_db struct {
	Db_name string
}

func (data_src *Json_db) Save(dict *map[int]int) {
	bytes, _ := json.Marshal(dict)
	os.WriteFile(data_src.Db_name, bytes, 0644)
}

func (data_src *Json_db) Retrieve(dict *map[int]int) {
	bytes, _ := os.ReadFile(data_src.Db_name)
	json.Unmarshal(bytes, &dict)
}
