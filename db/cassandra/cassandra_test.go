package cassandra

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ssolkhon/cf-keystore/cf"
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
	mySession, err := GetSession(myServices.Cassandra[0])
	if err != nil {
		t.Error("Error:", err)
	}
	defer mySession.Close()
}

func TestCreateTable(t *testing.T) {
	t.Log("Loading test config")
	myServices, err := loadTestData()
	if err != nil {
		t.Error("Error:", err)
	}
	t.Log("Getting session")
	mySession, err := GetSession(myServices.Cassandra[0])

	n := 10
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		t.Error("Error:", err)
	}
	myTableName := strings.ToLower(fmt.Sprintf("%X", b))

	t.Log("Generated random string", myTableName)
	t.Log("Creating table", myTableName)
	err = CreateTable(mySession, myServices.Cassandra[0].Credentials.Keyspace, myTableName)
	if err != nil {
		t.Error("Error:", err)
	}

	if result := tableExists(mySession, myServices.Cassandra[0].Credentials.Keyspace, myTableName); result != true {
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
	mySession, err := GetSession(myServices.Cassandra[0])
	t.Log("Adding", len(myCases), "test rows")
	for k, v := range myCases {
		err := CreateRow(mySession, myServices.Cassandra[0].Credentials.Keyspace, myTableName, k, v)
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
	mySession, err := GetSession(myServices.Cassandra[0])

	t.Log("Finding", len(myCases), "test rows")
	for k, v := range myCases {
		result, err := GetRow(mySession, myServices.Cassandra[0].Credentials.Keyspace, myTableName, k)
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
