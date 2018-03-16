package hello

const spanish = "Spanish"
const french = "French"
const vietnamese = "Vietnamese"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const vietnameseHelloPrefix = "Xin chao, "

func hello(name string, language string) string {
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
	case vietnamese:
		prefix = vietnameseHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}
