package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	yoruba             = "Yoruba"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	yorubaHelloPrefix  = "Bawoni, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case yoruba:
		prefix = yorubaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {

	fmt.Println(Hello("world", "Yoruba"))
}
