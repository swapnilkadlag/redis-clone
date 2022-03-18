package main

const (
	get = "get"
)

func getCommand(args []any) string {
	key := args[1]
	value := getKey(key)
	if i, isInt := value.(int); isInt {
		return integerStringOf(i)
	}
	if s, isString := value.(string); isString {
		return simpleStringOf(s)
	}
	return errorStringOf("Key not found")
}
