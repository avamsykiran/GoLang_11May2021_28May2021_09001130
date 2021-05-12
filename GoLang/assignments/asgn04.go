package main

import "fmt"

/*
  while caliculating simple interest,
    each bank can have different interest rates, but for
    a specific bank the interest rate is standard.

  create a function that accepts the rateOfInterest returns another funtion to compute simple interest
  given p and timePeriod
*/

func simpleInterest(r float64) func(float64, float64) float64 {
	return func(p float64, t float64) float64 {
		return (p * t * r) / 100
	}
}

func main() {
	hdfc := simpleInterest(8)
	sbi := simpleInterest(10.5)

	fmt.Printf("hdfc interest = %f \nsbi Interest = %f", hdfc(1000, 5), sbi(1000, 5))
}
