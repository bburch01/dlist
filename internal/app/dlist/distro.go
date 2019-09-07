package dlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Map of distro name to dependent module list
var distroDependencyMap = make(map[string][]string)

// Map of module name to distro
var moduleDistroMap = make(map[string]interface{})

func GetDepList(distro string) (map[string]string, error) {

	if err := initModuleDistroMap(); err != nil {
		return nil, err
	}

	if err := initDistroDependencyMap(); err != nil {
		panic(err)
	}
	if _, ok := distroDependencyMap[distro]; !ok {
		return nil, fmt.Errorf("%v is an invalid CPAN perl distribution name", distro)
	}

	depDistros := make(map[string]string)

	if err := resolveDependencies(distro, depDistros); err != nil {
		return nil, fmt.Errorf("dependency resolution failed with error: %v", err)
	}

	return depDistros, nil

}

func initDistroDependencyMap() error {

	err := filepath.Walk("data", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() != "data" {

				var distroMetaData map[string]interface{}

				jsonFile, err := os.Open(filepath.Join("data", info.Name(), "META.json"))

				if err != nil {
					return err
				}

				defer jsonFile.Close()

				jsonBytes, err := ioutil.ReadAll(jsonFile)
				if err != nil {
					return err

				}

				if err := json.Unmarshal(jsonBytes, &distroMetaData); err != nil {
					return err
				}

				distroDependencyMap[info.Name()] = extractDistroDeps(info.Name(), distroMetaData)

			}
		}

		return nil
	})

	return err
}

func initModuleDistroMap() error {

	jsonFile, err := os.Open("data/module-distro-map.json")

	if err != nil {
		return err

	}

	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err

	}

	if err := json.Unmarshal(jsonBytes, &moduleDistroMap); err != nil {
		return err
	}

	return nil
}

func extractDistroDeps(distro string, data map[string]interface{}) (deps []string) {

	var requires map[string]interface{}

	if _, ok := data["prereqs"]; ok {
		prereqs := data["prereqs"].(map[string]interface{})
		if _, ok := prereqs["runtime"]; ok {
			runtime := prereqs["runtime"].(map[string]interface{})
			if _, ok := runtime["requires"]; ok {
				requires = runtime["requires"].(map[string]interface{})
			}
		}
	}

	for k := range requires {
		if k != "perl" {
			if _, ok := CoreModulesMap[k]; !ok {
				deps = append(deps, k)
			}
		}
	}

	return
}

func resolveDependencies(distro string, depDistros map[string]string) error {

	if _, ok := depDistros[distro]; !ok {
		depDistros[distro] = distro
	}

	for _, v := range distroDependencyMap[distro] {
		resolveDependencies(moduleDistroMap[v].(string), depDistros)
	}

	return nil
}
