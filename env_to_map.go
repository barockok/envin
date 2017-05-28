package main

import (
	"regexp"
	"strconv"
	"strings"
)

func envToMap(currentMap map[string]interface{}, envKey string, value string) map[string]interface{} {
	if strings.Contains(envKey, "___") {
		keys := strings.Split(envKey, "___")
		if len(keys) == 1 {
			nativeTypeAssign(currentMap, envKey, value)
		} else {
			key := keys[0]
			if match, _ := regexp.MatchString("^[0-9]+$", key); match {
				// fmt.Println("*********** is integer key *********")
				// currentMap[key] = envToMap2Array(currentMap, key, envKey, value)
			} else {
				currentMap[key] = envToMap2Map(currentMap, key, envKey, value)
			}
		}
	} else {
		nativeTypeAssign(currentMap, envKey, value)
	}
	return currentMap
}

// func envToMap2Array(currentMap map[string]interface{}, key string, envKey string, value string) []interface{} {
// 	var cSlice []interface{}
// 	if currentMap[envKey] == nil {
// 		cSlice = make([]interface{}, 1)
// 	} else {
// 		cSlice = currentMap[envKey].([]interface{})
// 	}

// }

func envToMap2Map(currentMap map[string]interface{}, key string, envKey string, value string) map[string]interface{} {
	rootKey := key + "___"
	strippedKey := strings.Replace(envKey, rootKey, "", -1)
	var targetChild map[string]interface{}
	if currentMap[key] == nil {
		targetChild = make(map[string]interface{})
	} else {
		targetChild = currentMap[key].(map[string]interface{})
	}
	return envToMap(targetChild, strippedKey, value)
}

func nativeTypeAssign(targetmap map[string]interface{}, key string, value string) {

	var v interface{}
	if r, _ := regexp.MatchString("^-?[0-9]+\\.[0-9]+$", value); r {
		vFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			v = value
		} else {
			v = vFloat
		}
	} else if r, _ := regexp.MatchString("^-?[0-9]+$", value); r {
		vInt, err := strconv.Atoi(value)
		if err != nil {
			v = value
		} else {
			v = vInt
		}
	} else if value == "true" || value == "false" {
		vBool, err := strconv.ParseBool(value)
		if err != nil {
			v = value
		} else {
			v = vBool
		}
	} else {
		v = value
	}

	targetmap[key] = v
}
