package dlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// coreModulesList is an array containing the names of all of the modules shipped with the perl core.
var coreModulesList []string

// CoreModulesMap is a conversion of coreModulesList array to a map. It provides client code with fast look-ups to check
// for set membership.
var CoreModulesMap map[string]string

func init() {

	jsonFile, err := os.Open("data/core-modules.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &coreModulesList)

	CoreModulesMap = make(map[string]string)

	for _, s := range coreModulesList {
		CoreModulesMap[s] = s
	}

}
