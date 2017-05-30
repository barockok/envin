package main

import (
	"encoding/json"
	"testing"
)

func Test_mergeMap(t *testing.T) {
	dst := mapStringIface{"mysql_host": toI("localhost")}
	src := mapStringIface{"mysql_host": toI("127.0.0.1")}
	mergeMap(dst, src)
	if dst["mysql_host"] != "127.0.0.1" {
		t.Errorf("Error, failed overide value")
	}

	dst["redis_config"] = toI(mapStringIface{"time_out": toI(1000)})
	src["redis_config"] = toI(mapStringIface{"time_out": toI(1500)})
	mergeMap(dst, src)
	if dst["redis_config"].(mapStringIface)["time_out"].(int) != 1500 {
		t.Errorf("Error, failed overide value in deep map")
	}
	dst["rules"] = toI([]interface{}{toI("ACCEPT")})
	src["rules"] = toI(arrayCollector{entries: mapStringIface{"0": toI("CHALLENGE")}})
	mergeMap(dst, src)

	if dst["rules"].([]interface{})[0].(string) != "CHALLENGE" {
		t.Errorf("Error, simple array not parsed")
	}

}

func Test_mergeDestFromJson(t *testing.T) {
	jsonString := []byte(`
		{
			"simpleString" :  "this is string",
			"nestedAttr" : {
				"attrLevel1" : "hello"
			},
			"objectInArray" : [
				{
					"name" : "I'm so deep",
					"label" : "deep"
				}
			],
			"simpleArray" : [1,2,3,4]
		}
	`)
	dest := mapStringIface{}
	json.Unmarshal(jsonString, &dest)

	src := mapStringIface{}
	src["simpleArray"] = toI(arrayCollector{entries: mapStringIface{"2": toI(80)}})
	overrideObjectInArray := toI(mapStringIface{"name": "Not so deep anymore"})
	src["objectInArray"] = toI(arrayCollector{entries: mapStringIface{"0": overrideObjectInArray}})

	mergeMap(dest, src)
	overrideObjectInArrayRes := dest["objectInArray"].([]interface{})[0].(mapStringIface)
	if overrideObjectInArrayRes["name"] != "Not so deep anymore" {
		t.Errorf("Error, not override object in an array")
	}

	if overrideObjectInArrayRes["label"] != "deep" {
		t.Errorf("Error, its remove the objec in array that not being updated")
	}

	itemInSimpleArray := dest["simpleArray"].([]interface{})[2].(int)
	if itemInSimpleArray != 80 {
		t.Errorf("Error, not override simple value in an array")
	}

	src["newAttr"] = toI("hello")
	mergeMap(dest, src)
	if dest["newAttr"] != toI("hello") {
		t.Errorf("Error, cannot add new attribute")
	}
}
