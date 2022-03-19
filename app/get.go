package main

const (
	get = "get"
)

func getCommand(args []any) string {
	if len(args) < 2 {
		return errorStringOf(invalidArgs)
	}
	key := args[1]
	value := getKey(key)
	if value == nil {
		return errorStringOf(keyNotFound)
	}
	if i, isInt := value.(int); isInt {
		return integerStringOf(i)
	}
	if s, isString := value.(string); isString {
		return simpleStringOf(s)
	}
	return errorStringOf(invalidArg)
}
