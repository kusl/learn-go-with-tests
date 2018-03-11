package main

import "fmt"

// Français For French speakers
const Français = "Français"

// Español For Spanish speakers
const Español = "Español"

const helloPrefix = "Hello, "
const españolHelloPrefix = "Hola, "
const françaisHelloPrefix = "Bonjour, "

// Hello this function says hello to kus
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := helloPrefix
	switch language {
	case Français:
		prefix = françaisHelloPrefix
	case Español:
		prefix = españolHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("Kus", "English"))
}
