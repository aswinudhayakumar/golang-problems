package main

import (
	"fmt"
	"sync"
)

func printOdd(numberCh chan int, isDoneCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range numberCh {
		if n%2 != 0 {
			fmt.Println(n)
			isDoneCh <- true
		} else {
			numberCh <- n
		}
	}
}

func printEven(numberCh chan int, isDoneCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range numberCh {
		if n%2 == 0 {
			fmt.Println(n)
			isDoneCh <- true
		} else {
			numberCh <- n
		}
	}
}

func main() {
	var wg sync.WaitGroup
	numberCh := make(chan int)
	isDoneCh := make(chan bool)

	wg.Add(2)

	go printOdd(numberCh, isDoneCh, &wg)
	go printEven(numberCh, isDoneCh, &wg)

	for i := 0; i <= 10; i++ {
		numberCh <- i
		<-isDoneCh
	}

	close(numberCh)
	close(isDoneCh)

	wg.Wait()
}
