package main

import "fmt"

func main() {
	f()
}

func f() {
	fmt.Println("F called ")
	defer fmt.Println("This statement is defered")
	fmt.Println("F is complete")
}
