package main

import "fmt"

func main (){

	var x int = 56;

	//check if x is prime

	var count int = 2;

	for i:=2;i<=x/2;i++ {
		if(x%i==0){
			count++;
			break;
		}
	}

	if(count==2){
		fmt.Println("Its a prime")
	}else{
		fmt.Println("Its not a prime")
	}
}