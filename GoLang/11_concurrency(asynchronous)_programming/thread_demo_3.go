package main

import "fmt"
import "time"

func generateSeries(threadName string, start int, end int, isReverse bool) {
	if !isReverse {
		for i := start; i <= end; i++ {
			fmt.Println(threadName, ">> ", i)
			time.Sleep(time.Second)
		}
	} else {
		for i := start; i >= end; i-- {
			fmt.Println(threadName, ">> ", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	go generateSeries("thd1", 1, 10, false)
	go generateSeries("thd2", 10, 1, true)
	time.Sleep(10 * time.Second)
}
