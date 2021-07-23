package actions

func Help (msg string) string {
	msgNew := "You can choose the following commands: Hi, Bye"
	return msgNew
}

func Hi (msg string) string {
	msgNew := "Hello my friend!"
	return msgNew
}

func Bye (msg string) string {
	msgNew := "Good bye my friend"
	return msgNew
}

func Unknown () string {
	return "Unknown command, sorry. Please try again."
}
