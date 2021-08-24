package hello

import "fmt"

const englishHiPrefix = "Hi"
const exclamationSuffix = "!"
const spanishHiPrefix = "Hola"
const frenchHiPrefix = "Bonjour"

func sayHello(name string, language string) string {
	formattedName := ""
	if name != "" {
		formattedName = " " + name
	}

	return getGreetingPrefix(language) + formattedName + exclamationSuffix
}

func getGreetingPrefix(language string) (greetingPrefix string) {

	switch language {
	case "Spanish":
		greetingPrefix = spanishHiPrefix
	case "French":
		greetingPrefix = frenchHiPrefix
	default:
		greetingPrefix = englishHiPrefix
	}

	return
}

func main() {
	fmt.Println(sayHello("Siddiq", "English"))
}
