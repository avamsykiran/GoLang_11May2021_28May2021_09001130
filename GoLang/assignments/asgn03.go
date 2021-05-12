package main

import "fmt"

/*
   create a fucntion simpleInterest
   params:     p,t,r
   return:     si=(p*t*r)/100,amountToBePaid=p+si
*/

func simpleIntrest(principalAmount float64, rateOfIntrest float64, years float64) (float64, float64) {
	var si float64 = 0
	var totalAmount float64 = 0
	si = (principalAmount * rateOfIntrest * years) / 100
	totalAmount = principalAmount + si
	return si, totalAmount
}

func main() {
	fmt.Print("Result - ")
	fmt.Print(simpleIntrest(10000, 5.5, 2))
}
