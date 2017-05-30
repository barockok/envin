package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	conf := cliConf{}
	flag.BoolVar(&conf.backup, "backup", false, "backup the source file")
	flag.BoolVar(&conf.override, "override", false, "override the source file")
	flag.StringVar(&conf.kind, "filetype", "JSON", "file type of the source file, [JSON|YAML]")
	flag.StringVar(&conf.prefix, "prefix", "", "ENV prefix will used")
	flag.StringVar(&conf.srcfile, "src", "", "source file")
	flag.StringVar(&conf.targetfile, "target", "", "out file location")
	flag.Parse()

	if !conf.validate() {
		os.Exit(1)
	}

	customMap := make(mapStringIface)
	for key, value := range envMap(conf.prefix) {
		envToMap(customMap, key, value)
	}

	var loadedMap mapStringIface
	var err error

	if conf.srcfile != "" {
		if conf.kind == "JSON" {
			loadedMap, err = loadMapJSONFile(conf.srcfile)
		} else if conf.kind == "YAML" {
			loadedMap, err = loadMapYAMLFile(conf.srcfile)
		}
	} else {
		loadedMap = make(mapStringIface)
	}

	checkErrorAndExit(err)

	newMap := mergeMap(loadedMap, customMap)

	var newContent []byte
	if conf.kind == "JSON" {
		newContent, err = json.Marshal(newMap)
	} else {
		newContent, err = yaml.Marshal(newMap)
	}
	checkErrorAndExit(err)
	ioutil.WriteFile(conf.targetfile, newContent, 0644)
}

func envMap(prefix string) map[string]string {
	env := make(map[string]string)
	for _, i := range os.Environ() {
		if match, _ := regexp.MatchString("^"+prefix, i); match {
			ii := strings.Replace(i, prefix+"_", "", -1)
			sep := strings.Index(ii, "=")
			env[ii[0:sep]] = ii[sep+1:]
		} else {
			continue
		}
	}
	return env
}

func checkErrorAndExit(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error : %s", err))
		os.Exit(1)
	}
}
