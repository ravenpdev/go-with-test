package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// private function in go start with lowercase letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("raven", "Spanish"))
}
