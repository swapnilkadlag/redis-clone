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
	str, isString := args[1].(string)
	if !isString {
		return errorStringOf("Error")
	}
	return simpleStringOf(strings.ToUpper(str))
}
