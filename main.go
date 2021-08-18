package main

import "fmt"

func main() {
	fmt.Println(hello("world"))
}

const englishHelloPrefix = "hello, "

func hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}
