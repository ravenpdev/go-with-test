package main

const (
	spanish = "Spanish"
	french  = "French"

	englishHeloPrefix  = "Hello, "
	spanishHelloPrefix = "Holla, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixGreeting(language) + name
}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHeloPrefix
	}
	return
}
