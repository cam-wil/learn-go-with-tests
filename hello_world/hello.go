package main

import "fmt"

const EN_HELLO_PREFIX = "hello, "
const SP_HELLO_PREFIX = "hola, "
const FR_HELLO_PREFIX = "bonjour, "

const SPANISH = "spanish"
const FRENCH = "french"

func greetingPrefix(lang string) (prefix string) {
	switch lang {

	case SPANISH:
		prefix = SP_HELLO_PREFIX

	case FRENCH:
		prefix = FR_HELLO_PREFIX
	default:
		prefix = EN_HELLO_PREFIX
	}
	return
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("camer", ""))
}
