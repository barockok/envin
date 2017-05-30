package main

import (
	"fmt"
	"regexp"
)

type cliConf struct {
	kind       string
	override   bool
	srcfile    string
	targetfile string
	backup     bool
	prefix     string
}

func (c *cliConf) validate() bool {
	if c.targetfile == "" && c.srcfile != "" {
		c.targetfile = c.srcfile
	}

	if c.targetfile == "" && c.srcfile == "" {
		fmt.Println("please provide target or src")
		return false
	}

	if c.kind == "" {
		fmt.Println("prefix are required")
		return false
	}
	if c.prefix == "" {
		fmt.Println("prefix are required")
		return false
	}

	if match, _ := regexp.MatchString("JSON|YAML", c.kind); !match {
		fmt.Println("Only support YAML & JSON for file type")
		return false
	}
	return true
}
