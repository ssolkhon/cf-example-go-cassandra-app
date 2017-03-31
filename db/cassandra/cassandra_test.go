package cassandra

import (
	"cf-example-go-cassandra-app/cf"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const (
	DEFAULT_CONFIG = "./test_data/cassandra.json"
)

func loadTestData() (*cf.Services, error) {
	result := &cf.Services{}

	file, err := ioutil.ReadFile(DEFAULT_CONFIG)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(file, &result)
	return result, err
}

func TestGetSession(t *testing.T) {
	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}

	t.Log("Getting session")
	mySession, err := GetSession(myServices)
	if err != nil {
		t.Error("Error:", err)
	}
	defer mySession.Close()
}

func TestTableExists(t *testing.T) {
	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Getting session")
	mySession, err := GetSession(myServices)

	myCases := make(map[string]bool, 2)
	myCases["doesnotexist"] = false
	myCases["exists"] = true

	for k, e := range myCases {
		result := tableExists(mySession, k)
		if e != result {
			t.Error("Error: expected", e, "got", result)
		}
	}
}

func TestCreateTable(t *testing.T) {
	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Getting session")
	mySession, err := GetSession(myServices)

	n := 10
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		t.Error("Error:", err)
	}
	myTableName := strings.ToLower(fmt.Sprintf("%X", b))

	t.Log("Generated random string", myTableName)
	t.Log("Creating table", myTableName)
	err = CreateTable(mySession, myTableName)
	if err != nil {
		t.Error("Error:", err)
	}

	if result := tableExists(mySession, myTableName); result != true {
		t.Error("Error:", myTableName, "was not created successfully")
	}
}

func TestCreateRow(t *testing.T) {
	myTableName := "create_row_test"
	myCases := make(map[string]string, 5)

	myCases["keyone"] = "valueone"
	myCases["key2"] = "value2"
	myCases["keyThree"] = "valueThree"
	myCases["Key4"] = "Value4"
	myCases["12345"] = "54321"

	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Getting session")
	mySession, err := GetSession(myServices)
	t.Log("Adding", len(myCases), "test rows")
	for k, v := range myCases {
		err := CreateRow(mySession, myTableName, k, v)
		if err != nil {
			t.Error("Error:", err)
		} else {
			t.Log("Added", k, "|", v)
		}
	}
}

func TestGetRow(t *testing.T) {
	myTableName := "get_row_test"
	myCases := make(map[string]string, 5)

	myCases["keyone"] = "valueone"
	myCases["key2"] = "value2"
	myCases["keyThree"] = "valueThree"
	myCases["Key4"] = "Value4"
	myCases["12345"] = "54321"

	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Getting session")
	mySession, err := GetSession(myServices)

	t.Log("Finding", len(myCases), "test rows")
	for k, v := range myCases {
		result, err := GetRow(mySession, myTableName, k)
		if err != nil {
			t.Error("Error:", err)
		}
		if v != result {
			t.Error("Error: expected", v, "got", result)
		} else {
			t.Log("Found", k, "|", v)
		}
	}
}
