package main

import (
	"reflect"
	"strconv"
)

func mergeMap(dest map[string]interface{}, src map[string]interface{}) map[string]interface{} {
	for key, val := range dest {
		targetSrc := src[key]
		if targetSrc == nil {
			continue
		} else {
			if targetSrcVal, ok := targetSrc.(map[string]interface{}); reflect.TypeOf(val) == reflect.TypeOf(targetSrc) && ok {
				dest[key] = mergeMap(val.(map[string]interface{}), targetSrcVal)
			} else if coll, ok := targetSrc.(arrayCollector); ok {
				dest[key] = mergeArray(val, coll)
			} else {
				dest[key] = targetSrc
			}
		}
	}
	for key, val := range src {
		destAttr := dest[key]
		// attribute not present in dest, so src will add it
		if destAttr == nil {
			if _, ok := val.(map[string]interface{}); ok {
				emptyDest := map[string]interface{}{}
				dest[key] = mergeMap(emptyDest, val.(map[string]interface{}))
			} else if coll, ok := val.(arrayCollector); ok {
				emptyArrayIface := []interface{}{}
				dest[key] = mergeArray(toI(emptyArrayIface), coll)
			} else {
				dest[key] = val
			}
		}
	}
	return dest
}

func mergeArray(dest interface{}, src arrayCollector) interface{} {
	destArray := dest.([]interface{})
	for keyIndex, val := range src.entries {
		i, _ := strconv.Atoi(keyIndex)
		destVal := destArray[i]
		srcVal, okMapStringIface := val.(map[string]interface{})
		destValVal, okDestValIface := destVal.(map[string]interface{})
		if okMapStringIface && okDestValIface {
			destArray[i] = mergeMap(destValVal, srcVal)
		} else {
			destArray[i] = val
		}
	}
	return destArray
}
