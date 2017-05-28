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
			rootKey := keys[0] + "___"
			strippedKey := strings.Replace(envKey, rootKey, "", -1)
			var targetChild map[string]interface{}
			if currentMap[keys[0]] == nil {
				targetChild = make(map[string]interface{})
			} else {
				targetChild = currentMap[keys[0]].(map[string]interface{})
			}
			currentMap[keys[0]] = envToMap(targetChild, strippedKey, value)
		}
	} else {
		nativeTypeAssign(currentMap, envKey, value)
	}
	return currentMap
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
