package main

import "fmt"
import "time"

func generateSeries(threadName string, start int, end int, isReverse bool) chan string {
	ch := make(chan string)

	go func() {
		if !isReverse {
			for i := start; i <= end; i++ {
				ch <- fmt.Sprint(threadName, ">> ", i)
				time.Sleep(time.Second)
			}
		} else {
			for i := start; i >= end; i-- {
				ch <- fmt.Sprint(threadName, ">> ", i)
				time.Sleep(time.Second)
			}
		}
		close(ch)
	}()
	return ch
}

func printer(ch chan string) chan struct{} {
	signal := make(chan struct{})

	go func() {
		for data := range ch {
			fmt.Println(data)
			
			time.Sleep(time.Second)
		}
		close(signal)
	}()

	return signal
}

func main() {
	thd1 := generateSeries("thd1", 1, 10, false)
	thd2 := generateSeries("thd2", 10, 1, true)

	sig1 := printer(thd1)
	sig2 := printer(thd2)

	<-sig1
	<-sig2
}
