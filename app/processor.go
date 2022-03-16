package main

import (
	"fmt"
	"reflect"
)

func process(args []any) string {
	cmdObject := args[0]
	if reflect.ValueOf(cmdObject).Kind() == reflect.String {
		cmd := fmt.Sprint(args[0])
		switch cmd {
		case ping:
			pingCommand(args)
		case echo:
			echoCommand(args)
		}

		return pingCommand(args)
	}
	return "ERR"
}
