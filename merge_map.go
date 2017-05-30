package main

import (
	"reflect"
	"strconv"
)

func mergeMap(dest mapStringIface, src mapStringIface) mapStringIface {
	for key, val := range dest {
		targetSrc := src[key]
		if targetSrc == nil {
			continue
		} else {
			if targetSrcVal, ok := targetSrc.(mapStringIface); reflect.TypeOf(val) == reflect.TypeOf(targetSrc) && ok {
				dest[key] = mergeMap(val.(mapStringIface), targetSrcVal)
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
			if _, ok := val.(mapStringIface); ok {
				emptyDest := mapStringIface{}
				dest[key] = mergeMap(emptyDest, val.(mapStringIface))
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
		srcVal, okMapStringIface := val.(mapStringIface)
		destValVal, okDestValIface := destVal.(map[string]interface{})
		if okMapStringIface && okDestValIface {
			destArray[i] = mergeMap(destValVal, srcVal)
		} else {
			destArray[i] = val
		}
	}
	return destArray
}
