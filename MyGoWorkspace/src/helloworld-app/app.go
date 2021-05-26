package main

import (
	"fmt"
	"helloworld-app/greet"
)

func main() {
	userName := "Vamsy Kiran"

	fmt.Println(greet.Welcome(userName))
	fmt.Println(greet.SayBye(userName))
}
