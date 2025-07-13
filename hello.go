package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Elodie, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name + "!"
	}
	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
