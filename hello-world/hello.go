package main

import "fmt"

const spanish = "Spanish"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

// Hello returns "Hello, world" string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
