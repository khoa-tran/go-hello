package main

import "fmt"

const helloPrefix = "Hello, "

// Hello returns "Hello, world" string
func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
