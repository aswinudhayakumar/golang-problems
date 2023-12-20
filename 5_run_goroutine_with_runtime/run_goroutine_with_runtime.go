package main

import (
	"fmt"
)

func say(value string) {
	for i := 0; i < 5; i++ {
		// when Gosched() was called here,
		// from the normal function call, then only the go routine in line number #17 will run
		// runtime.Gosched()
		fmt.Println(value)
	}
}

func main() {
	go say("world")
	say("hello")
}

// ----------- output with runtime.Gosched() -----------
// hello
// world
// hello
// world
// hello
// world
// hello
// world
// hello

// ----------- output without runtime.Gosched() -----------
// hello
// hello
// hello
// hello
// hello
