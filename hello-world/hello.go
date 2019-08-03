package main

import "fmt"

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name
}

func main() {
	fmt.Println(hello("Ryan"))
	fmt.Println(hello(""))
}
