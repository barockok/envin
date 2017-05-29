package main

import (
	"testing"
)

func Test_mergeMap(t *testing.T) {
	dst := map[string]interface{}{"mysql_host": toI("localhost")}
	src := map[string]interface{}{"mysql_host": toI("127.0.0.1")}
	mergeMap(dst, src)
	if dst["mysql_host"] != "127.0.0.1" {
		t.Errorf("Error, failed overide value")
	}

	dst["redis_config"] = toI(map[string]interface{}{"time_out": toI(1000)})
	src["redis_config"] = toI(map[string]interface{}{"time_out": toI(1500)})
	mergeMap(dst, src)
	if dst["redis_config"].(map[string]interface{})["time_out"].(int) != 1500 {
		t.Errorf("Error, failed overide value in deep map")
	}
	dst["rules"] = toI([]interface{}{toI("ACCEPT")})
	src["rules"] = toI(arrayCollector{entries: map[string]interface{}{"0": toI("CHALLENGE")}})
	mergeMap(dst, src)

	if dst["rules"].([]interface{})[0].(string) != "CHALLENGE" {
		t.Errorf("Error, simple array not parsed")
	}

}
