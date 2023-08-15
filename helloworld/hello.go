package main

import "fmt"

// It is traditional for your first program in a new language to be Hello, World.

// hello.go

// func main() {
// 	fmt.Println("Hello World")
// }

//  Go automatically call main() function, so there is no need to call main() function explicitly and every executable program must contain ...

// We have created a new function again with func but this time we've added another
// keyword string in the definition.
// This means this function returns a string.

// func Hello(name string) string {
// 	return "Hello," + name
// }

// Constants are defined like so
const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchshHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
	//prefix := englishHelloPrefix

	//switch language {
	//case "French":
	//	prefix = frenchshHelloPrefix
	//case "Spanish":
	//	prefix = spanishHelloPrefix
	//}

	//return prefix + name

	// if language == "spanish" {
	// 	return spanishHelloPrefix + name
	// }

	// if language == "French" {
	// 	return frenchshHelloPrefix + name
	// }

	// return englishHelloPrefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchshHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {
	fmt.Println(Hello("Elodie", "Spanish"))
}
