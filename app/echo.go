package main

import (
	"strings"
)

const (
	echo = "echo"
)

func echoCommand(args []any) string {
	if len(args) < 2 {
		return errorStringOf(invalidArgs)
	}
	if s, isString := args[1].(string); isString {
		return simpleStringOf(strings.ToUpper(s))
	}
	if i, isInt := args[1].(int); isInt {
		return integerStringOf(i)
	}
	return errorStringOf(invalidArg)
}
