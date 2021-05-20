package main

import (
	"fmt"
	"time"
)

func generateSeries(start int, end int) <-chan int {
	ch := make(chan int)

	go func() {

		for i := start; i <= end; i++ {
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()
	return ch
}

func main() {
	thd1 := generateSeries(1, 10)
	thd2 := generateSeries(100, 110)

	for {
		select {
		case data1 := <-thd1:
			fmt.Println("From First Thread>> ", data1)
		case data2 := <-thd2:
			fmt.Println("From Second Thread>> ", data2)
			/*case <-time.After(2 * time.Second):
			fmt.Println("Termianted Because of time out")
			break */
		}
	}
}
