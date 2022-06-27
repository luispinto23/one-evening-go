package main

import "fmt"

func main() {
	Greet("Alice")
	Greet("Bob")
}

func Greet(name string) {
	fmt.Printf("Hello, %s\n", name)
}
