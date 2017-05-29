package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func loadMapJSONFile(filepath string) (map[string]interface{}, error) {
	var err error
	if _, err := os.Stat(filepath); err != nil {
		return nil, &readFileError{filepath, "JSON"}
	}
	jsonBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, &readFileError{filepath, "JSON"}
	}
	predifinedMap := map[string]interface{}{}
	err = json.Unmarshal(jsonBytes, &predifinedMap)
	if err != nil {
		return nil, &readFileError{filepath, "JSON"}
	}

	return predifinedMap, nil
}

func loadMapYAMLFile(filepath string) (map[string]interface{}, error) {
	var err error
	if _, err := os.Stat(filepath); err != nil {
		return nil, &readFileError{filepath, "YAML"}
	}
	yamlBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, &readFileError{filepath, "YAML"}
	}
	predifinedMap := map[string]interface{}{}
	err = yaml.Unmarshal(yamlBytes, &predifinedMap)
	if err != nil {
		return nil, &readFileError{filepath, "YAML"}
	}
	return predifinedMap, nil
}
