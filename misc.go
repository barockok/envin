package main

import "fmt"

type readFileError struct {
	path string
	kind string
}

func (instance *readFileError) Error() string {
	return fmt.Sprintf("cannot read file in %s with type %s", instance.path, instance.kind)
}
