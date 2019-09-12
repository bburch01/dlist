package dlist

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var coreModulesList []string
var CoreModulesMap map[string]string

func init() {

	if err := godotenv.Load(); err != nil {
		log.Panicf("failed to load environment variables with error: %v", err)
	}

	jsonFile, err := os.Open(os.Getenv("DATA_DIR") + "/core-modules.json")
	if err != nil {
		log.Panicf("failed to load core-modules.json with error: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &coreModulesList)

	CoreModulesMap = make(map[string]string)
	
	for _, s := range coreModulesList {
		CoreModulesMap[s] = s
	}
}
