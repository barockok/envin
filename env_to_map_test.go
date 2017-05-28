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

	// var arrAttr []interface{}
	// envToMap(predefinedMap, "items___0___name", "orange")
	// envToMap(predefinedMap, "items___0___qty", "1")

	// envToMap(predefinedMap, "items___1___name", "mango")
	// envToMap(predefinedMap, "items___1___qty", "2")

	// if tipe, isTrue := predefinedMap["items"].([]interface{}); !isTrue {
	// 	t.Errorf("Error, items not initialized as array type : %v ", tipe)
	// }
	// arrAttr = predefinedMap["items"].([]interface{})

	// if len(arrAttr) != 2 {
	// 	t.Errorf("Error, arr attr not parsed")
	// }
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
