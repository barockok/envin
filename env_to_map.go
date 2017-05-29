package main

import (
	"regexp"
	"strconv"
	"strings"
)

type arrayCollector struct {
	entries map[string]interface{}
}

func envToMap(currentMap map[string]interface{}, envKey string, value string) map[string]interface{} {
	if strings.Contains(envKey, "___") {
		keys := strings.Split(envKey, "___")
		if len(keys) == 1 {
			nativeTypeAssign(currentMap, envKey, value)
		} else {
			key := keys[0]
			key2 := keys[1]
			key1Match, _ := regexp.MatchString("^[0-9]+$", key)
			key2Match, _ := regexp.MatchString("^[0-9]+$", key2)

			if !key1Match && key2Match {
				currentMap[key] = envToMap2Array(currentMap, envKey, key, key2, value)
			} else {
				currentMap[key] = envToMap2Map(currentMap, key, envKey, value)
			}
		}
	} else {
		nativeTypeAssign(currentMap, envKey, value)
	}
	return currentMap
}

func envToMap2Array(cmap map[string]interface{}, envKey, collKey, itemKey, value string) interface{} {
	var coll arrayCollector
	if cmap[collKey] == nil {
		coll = arrayCollector{}
		coll.entries = map[string]interface{}{}
	} else {
		coll = cmap[collKey].(arrayCollector)
	}

	rootKey := collKey + "___" + itemKey
	strippedKey := strings.Replace(envKey, rootKey, "", -1)

	if strippedKey == "" {
		coll.entries[itemKey] = convString(value)
	} else {
		rootKey := collKey + "___" + itemKey + "___"
		strippedKey := strings.Replace(envKey, rootKey, "", -1)

		var itemCont map[string]interface{}
		if itemContIface := coll.entries[itemKey]; itemContIface == nil {
			itemCont = map[string]interface{}{}
		} else {
			itemCont = itemContIface.(map[string]interface{})
		}
		coll.entries[itemKey] = envToMap(itemCont, strippedKey, value)
	}
	return coll
}

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
	targetmap[key] = convString(value)
}

func convString(value string) interface{} {
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
	return v
}
