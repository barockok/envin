package main

import (
	"testing"
)

func Test_envToMap(t *testing.T) {
	predefinedMap := make(map[string]interface{})

	envToMap(predefinedMap, "mysql_host", "localhost")
	if predefinedMap["mysql_host"] != "localhost" {
		t.Errorf("failed assign simple value expected %s, got %s", "localhost", predefinedMap["mysql_host"])
	}

	envToMap(predefinedMap, "mysql_host", "127.0.0.1")
	if predefinedMap["mysql_host"] == "localhost" {
		t.Errorf("failed override simple value expected %s, got %s", "127.0.0.1", predefinedMap["mysql_host"])
	}

	var mysqlMap map[string]interface{}
	envToMap(predefinedMap, "mysql___host", "127.0.0.1")
	if tipe, ok := predefinedMap["mysql"].(map[string]interface{}); !ok {
		t.Errorf("error, nested key not converted to map interface , got %s", tipe)
	}

	mysqlMap = predefinedMap["mysql"].(map[string]interface{})
	if mysqlMap["host"] != "127.0.0.1" {
		t.Errorf("error, nested key not assigned correctly")
	}

	envToMap(predefinedMap, "mysql___port", "\"3306\"")
	mysqlMap = predefinedMap["mysql"].(map[string]interface{})
	if !(mysqlMap["host"] == "127.0.0.1" && mysqlMap["port"] == "\"3306\"") {
		t.Errorf("error, add new item on preexisting nested map failed %v", mysqlMap["host"])
	}

	envToMap(predefinedMap, "items___0___name", "orange")
	envToMap(predefinedMap, "items___0___qty", "1")

	if tipe, ok := predefinedMap["items"].(arrayCollector); !ok {
		t.Errorf("Error, items not initialized as arrayCollectore, go type : %v ", tipe)
	}

	item0 := predefinedMap["items"].(arrayCollector).entries["0"].(map[string]interface{})
	if item0["name"] != "orange" {
		t.Errorf("Error, items content not parsed correctly expected : %s got %s", "orange", item0["name"])
	}
	if item0["qty"] != 1 {
		t.Errorf("Error, items content not parsed correctly expected : %d got %s", 1, item0["qty"])
	}
	envToMap(predefinedMap, "roles___0", "admin")
	envToMap(predefinedMap, "roles___1", "supreme")
	roles := predefinedMap["roles"].(arrayCollector).entries
	if actLen := len(roles); actLen != 2 {
		t.Errorf("Error, number of array parsed incorrect, expected %d, got : %d", 2, actLen)
	}

}

func Test_nativeTypeAssign(t *testing.T) {
	predefinedMap := make(map[string]interface{})
	nativeTypeAssign(predefinedMap, "anInteger", "100")
	if predefinedMap["anInteger"] != 100 {
		t.Errorf("Failed to format integer")
	}

	nativeTypeAssign(predefinedMap, "anInteger", "-100")
	if predefinedMap["anInteger"] != -100 {
		t.Errorf("Failed to format negative integer")
	}

	nativeTypeAssign(predefinedMap, "aFloat", "1.230")
	if predefinedMap["aFloat"] != 1.23 {
		t.Errorf("Failed to format Float ")
	}

	nativeTypeAssign(predefinedMap, "aFloat", "-1.230")
	if predefinedMap["aFloat"] != -1.23 {
		t.Errorf("Failed to format Negative Float")
	}

	nativeTypeAssign(predefinedMap, "aBool", "true")
	if predefinedMap["aBool"] != true {
		t.Errorf("Failed to format boolean true")
	}

	nativeTypeAssign(predefinedMap, "aBool", "false")
	if predefinedMap["aBool"] != false {
		t.Errorf("Failed to format boolean false")
	}

}
