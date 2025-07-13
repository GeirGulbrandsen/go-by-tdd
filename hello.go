package main

import "fmt"

const defaultEnglishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Elodie, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name + "!"
	}
	return defaultEnglishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
