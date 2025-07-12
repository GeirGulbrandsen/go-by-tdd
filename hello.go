package main

import "fmt"

func Hello(s string) string {
	const englishHelloPrefix = "Hello, "
	return englishHelloPrefix + s + "!"
}

func main() {
	fmt.Println(Hello("world"))
}
