package main

import "fmt"

func main() {
	mainCh := make(chan int)
	// using struct{} for quit channel will reduce memory instead of using int
	quitCh := make(chan struct{})
	n := 5

	go func() {
		sum := 0
		for i := 0; i < n; i++ {
			sum += <-mainCh
		}
		fmt.Println(sum)
		quitCh <- struct{}{}
	}()

	SumOfSquares(mainCh, quitCh)
}

func SumOfSquares(mainCh chan int, quitCh chan struct{}) {
	y := 1
	for {
		select {
		case mainCh <- (y * y):
			y++
		case <-quitCh:
			return
		}
	}
}
