package dlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Result map[string]interface{}

func init() {

	jsonFile, err := os.Open("data/DateTime/META.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &Result)
}
