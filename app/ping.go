package main

import (
	"fmt"
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
	return simpleStringOf(strings.ToUpper(fmt.Sprint(args[1])))
}
