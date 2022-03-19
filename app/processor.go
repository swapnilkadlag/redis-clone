package main

func process(args []any) string {
	if cmd, isString := args[0].(string); isString {
		switch cmd {
		case ping:
			return pingCommand(args)
		case echo:
			return echoCommand(args)
		case set:
			return setCommand(args)
		case get:
			return getCommand(args)
		}
		return errorStringOf(invalidCommand)
	}
	return errorStringOf(invalidArg)
}
