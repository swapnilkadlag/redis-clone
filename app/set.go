package main

const (
	set = "set"
	px  = "px"
)

func setCommand(args []any) string {
	if len(args) < 3 {
		return errorStringOf(invalidArgs)
	}
	key := args[1]
	value := args[2]
	pxv := -1
	if index := indexOf(args, px); index > null {
		if i, isInt := args[index+1].(int); isInt {
			pxv = i
		} else {
			return errorStringOf(invalidArg)
		}
	}
	setKey(key, value, pxv)
	return simpleStringOf(ok)
}
