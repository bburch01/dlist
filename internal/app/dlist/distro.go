package dlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var distroDependencyMap = make(map[string][]string)
var moduleDistroMap = make(map[string]interface{})

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panicf("failed to load environment variables with error: %v", err)
	}
}

func GetDepList(distroList []string) (string, error) {
	var sb strings.Builder
	if err := initModuleDistroMap(); err != nil {
		return sb.String(), err
	}
	if err := initDistroDependencyMap(); err != nil {
		panic(err)
	}
	sb.WriteString("\n{\n")
	for _, v := range distroList {
		if _, ok := distroDependencyMap[v]; !ok {
			return sb.String(), fmt.Errorf("%v is an invalid CPAN perl distribution name", v)
		}
		sb.WriteString("\t\"")
		sb.WriteString(v)
		sb.WriteString("\": {")
		if len(distroDependencyMap[v]) == 0 {
			sb.WriteString("}")
		} else {
			if err := resolveDependencies(v, 0, &sb); err != nil {
				return sb.String(), fmt.Errorf("dependency resolution failed with error: %v", err)
			}
		}
	}
	sb.WriteString("\n}")
	jsonBytes := []byte(sb.String())
	scrubInvalidTrailingCommas(jsonBytes)
	return string(jsonBytes), nil
}

func initDistroDependencyMap() error {
	err := filepath.Walk(os.Getenv("DATA_DIR"), func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if info.Name() != "data" {
				var distroMetaData map[string]interface{}
				jsonFile, err := os.Open(os.Getenv("DATA_DIR") + "/" + info.Name() + "/META.json")
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
	jsonFile, err := os.Open(os.Getenv("DATA_DIR") + "/module-distro-map.json")
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

func resolveDependencies(distro string, level int, sb *strings.Builder) error {
	level++
	for i, v := range distroDependencyMap[distro] {
		sb.WriteString("\n")
		for t := 0; t < level+1; t++ {
			sb.WriteString("\t")
		}
		sb.WriteString("\"")
		sb.WriteString(moduleDistroMap[v].(string))
		if len(distroDependencyMap[moduleDistroMap[v].(string)]) > 0 {
			sb.WriteString("\": {")
		} else {
			sb.WriteString("\": {}")
			if i < len(distroDependencyMap[distro])-1 {
				sb.WriteString(",")
			}
		}
		resolveDependencies(moduleDistroMap[v].(string), level, sb)
	}
	if len(distroDependencyMap[distro]) > 0 {
		sb.WriteString("\n")
		for t := 0; t < level; t++ {
			sb.WriteString("\t")
		}
		sb.WriteString("},")
	}
	return nil
}

func scrubInvalidTrailingCommas(strBytes []byte) {
	prevCloseBraceExists := false
	prevCloseBraceIndex := 0
	patternMatched := false
	for i, v := range strBytes {
		if v == '}' {
			if prevCloseBraceExists {
				if strBytes[prevCloseBraceIndex+1] == ',' {
					patternMatched = true
					for idx := prevCloseBraceIndex + 2; idx < i; idx++ {
						if strBytes[idx] != '\n' && strBytes[idx] != '\t' {
							patternMatched = false
							break
						}
					}
					if patternMatched {
						strBytes[prevCloseBraceIndex+1] = ' '
					}
				}
				prevCloseBraceIndex = i
			} else {
				prevCloseBraceIndex = i
				prevCloseBraceExists = true
			}
		}
	}
	return
}
