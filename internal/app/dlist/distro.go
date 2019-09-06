package dlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var DistroMetaDataFilepathMap = make(map[string]string)

var DistroMetaDataMap = make(map[string]map[string]interface{})

func init() {

	//distroMetaDataFilepathMap := make(map[string]string)

	err := filepath.Walk("data", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {

			if info.Name() != "data" {

				DistroMetaDataFilepathMap[info.Name()] = filepath.Join("data", info.Name(), "META.json")

			}

		}

		//files = append(files, path)
		return nil
	})

	// TODO: add a proper error handler
	if err != nil {
		fmt.Println(err)
	}

	//DistroMetaDataMap = make(map[string]interface{})

	for k, v := range DistroMetaDataFilepathMap {

		var result map[string]interface{}

		//fmt.Printf("\nfilepath to open: %v", v)

		jsonFile, err := os.Open(v)

		// TODO: add proper error handler
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal([]byte(byteValue), &result)

		DistroMetaDataMap[k] = result

	}

	/*
		jsonFile, err := os.Open("data/DateTime/META.json")

		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal([]byte(byteValue), &Result)

	*/
}
