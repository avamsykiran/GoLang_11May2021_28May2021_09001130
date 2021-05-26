package main

import "fmt"

func sum(nums ...int) int {
	s := 0

	for _, val := range nums {
		s += val
	}

	return s
}

func main() {
	fmt.Printf("\nSum %d", sum(10, 20))
	fmt.Printf("\nSum %d", sum(10, 20, 30))
	fmt.Printf("\nSum %d", sum(10, 20, 30, 40, 50, 60))
}
