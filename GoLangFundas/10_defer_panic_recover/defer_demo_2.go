package main

import "fmt"

func half(input int) (output int) {
	defer func() {
		output++
	}()

	return input / 2
}

func main() {

	n1 := half(10)
	n2 := half(25)

	fmt.Println(n1, "\t", n2)
}
