package main

import (
	"reflect"
	"strconv"
)

func mergeMap(dest, src map[string]interface{}) map[string]interface{} {
	for key, val := range dest {
		targetSrc := src[key]
		if targetSrc == nil {
			continue
		} else {
			if _, ok := targetSrc.(map[string]interface{}); reflect.TypeOf(val) == reflect.TypeOf(targetSrc) && ok {
				dest[key] = mergeMap(val.(map[string]interface{}), targetSrc.(map[string]interface{}))
			} else if coll, ok := targetSrc.(arrayCollector); ok {
				dest[key] = mergeArray(val, coll)
			} else {
				dest[key] = targetSrc
			}
		}
	}
	return dest
}

func mergeArray(dest interface{}, src arrayCollector) interface{} {
	destArray := dest.([]interface{})
	for i, val := range destArray {
		targetSrc := src.entries[strconv.Itoa(i)]
		valVal, ok1 := val.(map[string]interface{})
		targetVal, ok2 := targetSrc.(map[string]interface{})

		if targetSrc == nil {
			continue
		} else if ok1 && ok2 {
			destArray[i] = mergeMap(valVal, targetVal)
		} else {
			destArray[i] = targetSrc
		}
	}
	return destArray
}
