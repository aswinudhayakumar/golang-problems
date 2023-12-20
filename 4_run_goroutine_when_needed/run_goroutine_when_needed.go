package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside the function")
	go func() {
		fmt.Println("Inside the function")
	}()
	fmt.Println("Oustide the function")

	runtime.Gosched()
}
