package main

import (
	"encoding/json"
	"os"
)

type Row struct {
	Alias string
	Link string
}

func marshalJson(row Row) ([]byte) {
	json, err := json.Marshal(row)
	if err != nil {
		panic(err)
	}
	return json
}

func unmarshalJson(jsonData []byte) Row {
	var bookmarkForRecovery Row
	err := json.Unmarshal(jsonData, &bookmarkForRecovery)
	if err != nil {
		panic(err)
	}
	return bookmarkForRecovery
}

func fileSaver(bookmarkJson []byte, err error) {
		err = os.WriteFile("user.json", bookmarkJson, 0644)
	if err != nil {
		panic(err)
	}
}

func fileReader() []byte {
	jsonData, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}
	return jsonData
}