package main

import "fmt"

//recursive fucntion called sumUntilN
//shall return the sum of first n natural numbers for a given n 
// sumUntilN(3) -> 1+2+3 = 6

/*func sumUntilN( n int)  int {
	return (n*(n+1))/2
}*/

func sumUntilN(n int) int {
    var result int = 0
    
	if n > 0 {
        result = n + sumUntilN(n-1)
    }

    return result
}

func main (){
	fmt.Println(sumUntilN(10))
}