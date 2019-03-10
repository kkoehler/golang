package main

import "fmt"

func Add(first int, second int) int {
	return first + second
}

func main() {
	fmt.Printf("Hello World, %d \n", Add(10, 10))
}
