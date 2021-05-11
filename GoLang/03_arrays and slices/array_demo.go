package main

import "fmt"

func main(){

	var nums = [5]int{10,20,30,40,50}

	for index,num := range nums {
		fmt.Println(index,"-Element: ",num)
	}
}