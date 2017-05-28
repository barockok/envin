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

	envToMap(predefinedMap, "mysql___port", "3306")
	mysqlMap = predefinedMap["mysql"].(map[string]interface{})
	if !(mysqlMap["host"] == "127.0.0.1" && mysqlMap["port"] == "3306") {
		t.Errorf("error, add new item on preexisting nested map failed %v", mysqlMap["host"])
	}
}
