package main

import (
	"fmt"
	"strings"
)

const (
	echo = "echo"
)

func echoCommand(args []any) string {
	return simpleStringOf(strings.ToUpper(fmt.Sprint(args[1])))
}
