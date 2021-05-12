package main

import "fmt"

func greetGenerator(input string) func(string) string {
	greetNote := input

	return func(unm string) string {
		return greetNote + unm
	}
}

func main() {
	greetUserInHindi := greetGenerator("Namsaskar ")
	greetUserInEnglish := greetGenerator("Hello ")
	greetUserInTamil := greetGenerator("Vanakkam ")

	fmt.Println(greetUserInHindi("Vamsy"))
	fmt.Println(greetUserInEnglish("Vamsy"))
	fmt.Println(greetUserInTamil("Vamsy"))
}
