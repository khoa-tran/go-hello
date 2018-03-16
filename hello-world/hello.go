package hello

const spanish = "Spanish"
const french = "French"
const vietnamese = "Vietnamese"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const vietnameseHelloPrefix = "Xin chao, "

// Hello says hello to specified name in a language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return GreetingPrefix(language) + name
}

// GreetingPrefix returns a hello prefix in a language
func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case vietnamese:
		prefix = vietnameseHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
