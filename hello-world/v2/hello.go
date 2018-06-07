package main

import "fmt"

// Hello returns a greeting
func Hello(name string, language string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Default name", "Default language"))
}
