package main

import "fmt"
import g "greet"

func main() {
	fmt.Println("This is a common workspace structure")

	fmt.Println("Auhtor: ", author)
	fmt.Println("Version: ", version)
	fmt.Println("Description: ", description)

	fmt.Println(g.GreetUser("Vamsy"))
}
