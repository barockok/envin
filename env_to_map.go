package main

import (
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
	var v interface{} = value
	targetmap[key] = v
}
