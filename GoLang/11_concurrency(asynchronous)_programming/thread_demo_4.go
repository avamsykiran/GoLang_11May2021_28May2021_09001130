package main

import "fmt"
import "time"

func generateSeries(threadName string, start int, end int, isReverse bool) <-chan struct{} {
	ch := make(chan struct{})

	go func() {
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
		close(ch)
	}()
	return ch
}

func main() {
	thd1 := generateSeries("thd1", 1, 10, false)
	thd2 := generateSeries("thd2", 10, 1, true)

	<-thd1
	<-thd2
}
