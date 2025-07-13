package main

import "fmt"

const spanish = "Spanish"
const defaultEnglishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Elodie, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == spanish {
		return spanishHelloPrefix + name + "!"
	}
	return defaultEnglishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
