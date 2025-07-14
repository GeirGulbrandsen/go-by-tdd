package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const norwegian = "Norwegian"

const defaultEnglishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const norwegianHelloPrefix = "Hei, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := defaultEnglishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case norwegian:
		prefix = norwegianHelloPrefix
	}
	return prefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
