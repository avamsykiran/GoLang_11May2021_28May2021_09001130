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


func main() {
	thd1 := generateSeries("thd1", 1, 10, false)
	thd2 := generateSeries("thd2", 10, 1, true)

	for{
		select {
		case data1 := <- thd1:
			fmt.Println(data1)
		case data2 := <- thd2:
			fmt.Println(data2)
		case <-time.After(2*time.Second):
			fmt.Println("Termianted Because of time out")
			break
		}
	}
}
