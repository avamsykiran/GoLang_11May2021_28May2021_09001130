package main

import "fmt"

func greetGenerator() func(string) string {
	return func(unm string) string {
		return "Hello " + unm
	}
}

func main() {

	greetUser := greetGenerator()

	fmt.Println(greetUser("Vamsy"))
}
