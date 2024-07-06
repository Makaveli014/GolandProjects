package main

import "fmt"

func main() {
	var a string
	var b string
	fmt.Scanln(&a, &b)
	fmt.Println("Привет", b)
	fmt.Println("Как тебя зовут?", a)
}
