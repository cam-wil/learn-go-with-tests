package main

import "fmt"

const HELLO_PREFIX = "hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return HELLO_PREFIX + name
}

func main() {
	fmt.Println(Hello("camer"))
}
