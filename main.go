package main

import "fmt"

func main() {
	fmt.Println(hello("Peter"))
}

func hello(name string) string {

	return "Hello" + name
}
