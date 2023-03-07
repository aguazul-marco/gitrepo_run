package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("5 + 5 =", add(5, 5))
	fmt.Println("5 - 5 =", sub(5, 5))
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
