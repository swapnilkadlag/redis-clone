package main

import (
	"strings"
)

const (
	ping = "ping"
	pong = "pong"
)

func pingCommand(args []any) string {
	if len(args) == 1 {
		return simpleStringOf(strings.ToUpper(pong))
	}
	if s, isString := args[1].(string); isString {
		return simpleStringOf(s)
	}
	if i, isInt := args[1].(int); isInt {
		return integerStringOf(i)
	}
	return errorStringOf(invalidArg)
}
