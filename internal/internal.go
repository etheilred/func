package internal

import (
	"fmt"
	"func/parser"
)

var variables map[string]int64 = make(map[string]int64)

// Set sets variable value and returns the value
func Set(id string, val int64) (parser.Attrib, error) {
	variables[id] = val
	return parser.Attrib(val), nil
}

// Get returns variable value, or an error if variable does not exist
func Get(id string) (parser.Attrib, error) {
	if _, ok := variables[id]; !ok {
		return 0, fmt.Errorf("No such vaiable: %s", id)
	}
	return parser.Attrib(variables[id]), nil
}
