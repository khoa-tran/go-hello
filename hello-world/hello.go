package main

import "fmt"

const helloPrefix = "Hello, "

// Hello returns "Hello, world" string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
