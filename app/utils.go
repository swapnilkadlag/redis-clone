package main

func contains(list []any, object any) int {
	for i := 0; i < len(list); i++ {
		if list[i] == object {
			return i
		}
	}
	return -1
}

func simpleStringOf(s string) string {
	return string(simpleString) + s + crlf
}

func bulkStringOf(s string) string {
	return string(bulkString) + s + crlf
}

func errorStringOf(s string) string {
	return string(err) + s + crlf
}

func integerStringOf(i int) string {
	return string(integer) + string(i)
}
