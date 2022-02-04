package main

import "fmt"

const (
	SPANISH = "Spanish"
	DEUTSCH = "Deutsch"
)
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const deutschHelloPrefix = "Hallo "

func getLanguagePrefix(language string) (prefix string) {
	switch language {
	case DEUTSCH:
		prefix = deutschHelloPrefix
	case SPANISH:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func sayHello(name, language string) string {
	if name == "" {
		return getLanguagePrefix(language) + "world"
	}
	return getLanguagePrefix(language) + name
}

func main() {
	fmt.Println(sayHello("to you", "English"))
}