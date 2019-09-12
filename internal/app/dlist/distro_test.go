package dlist

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Panicf("failed to load environment variables with error: %v", err)
	}
}

func TestGetDepListForSingleDistro(t *testing.T) {

	distroList := []string{"Class-Load"}
	var distroDeps map[string]interface{}

	if jsonStr, err := GetDepList(distroList); err == nil {
		jsonBytes := []byte(jsonStr)
		if err := json.Unmarshal(jsonBytes, &distroDeps); err != nil {
			t.Error(fmt.Printf("test failed with error %v", err))
			t.FailNow()
		}
	} else {
		t.Error("test failed with error: ", err)
		t.FailNow()
	}

	if _, ok := distroDeps["Class-Load"]; !ok {
		t.Error("Class-Load not found.")
		t.FailNow()
	}
	classLoad := distroDeps["Class-Load"].(map[string]interface{})

	if _, ok := classLoad["Try-Tiny"]; !ok {
		t.Error("Class-Load[Try-Tiny] not found.")
		t.FailNow()
	}
	if _, ok := classLoad["Module-Runtime"]; !ok {
		t.Error("Class-Load[Module-Runtime] not found.")
		t.FailNow()
	}

	if _, ok := classLoad["Data-OptList"]; !ok {
		t.Error("Class-Load[Data-OptList] not found.")
		t.FailNow()
	}
	dataOptList := classLoad["Data-OptList"].(map[string]interface{})

	if _, ok := dataOptList["Params-Util"]; !ok {
		t.Error("Class-Load[Data-OptList[Params-Util]] not found.")
		t.FailNow()
	}
	if _, ok := dataOptList["Sub-Install"]; !ok {
		t.Error("Class-Load[Data-OptList[Sub-Install]] not found.")
		t.FailNow()
	}

	if _, ok := classLoad["Module-Implementation"]; !ok {
		t.Error("Class-Load[Module-Implementation] not found.")
		t.FailNow()
	}
	moduleImplementation := classLoad["Module-Implementation"].(map[string]interface{})

	if _, ok := moduleImplementation["Module-Runtime"]; !ok {
		t.Error("Class-Load[Module-Implementation[Module-Runtime]] not found.")
		t.FailNow()
	}
	if _, ok := moduleImplementation["Try-Tiny"]; !ok {
		t.Error("Class-Load[Module-Implementation[Try-Tiny]] not found.")
		t.FailNow()
	}

	if _, ok := classLoad["Package-Stash"]; !ok {
		t.Error("Class-Load[Package-Stash] not found.")
		t.FailNow()
	}
	packageStash := classLoad["Package-Stash"].(map[string]interface{})

	if _, ok := packageStash["Module-Implementation"]; !ok {
		t.Error("Class-Load[Package-Stash[Module-Implementation]] not found.")
		t.FailNow()
	}
	moduleImplementation = packageStash["Module-Implementation"].(map[string]interface{})

	if _, ok := moduleImplementation["Module-Runtime"]; !ok {
		t.Error("Class-Load[Package-Stash[Module-Implementation[Module-Runtime]]] not found.")
		t.FailNow()
	}
	if _, ok := moduleImplementation["Try-Tiny"]; !ok {
		t.Error("Class-Load[Package-Stash[Module-Implementation[Try-Tiny]]] not found.")
		t.FailNow()
	}

	if _, ok := packageStash["Dist-CheckConflicts"]; !ok {
		t.Error("Class-Load[Package-Stash[Dist-CheckConflicts]] not found.")
		t.FailNow()
	}
	distCheckConflicts := packageStash["Dist-CheckConflicts"].(map[string]interface{})

	if _, ok := distCheckConflicts["Module-Runtime"]; !ok {
		t.Error("Class-Load[Package-Stash[Dist-CheckConflicts[Module-Runtime]]] not found.")
		t.FailNow()
	}
}
