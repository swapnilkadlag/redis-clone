package main

const (
	set = "set"
	px  = "px"
)

func setCommand(args []any) string {
	key := args[1]
	value := args[2]
	px := -1
	if len(args) > 3 {
		if i, isInt := args[4].(int); isInt {
			px = i
		}
	}
	setKey(key, value, px)
	return simpleStringOf(ok)
}
