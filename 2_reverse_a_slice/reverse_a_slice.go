package main

import "fmt"

func reverseASlice(data []int) []int {
	for a, b := 0, len(data)-1; a < b; a, b = a+1, b-1 {
		data[a], data[b] = data[b], data[a]
	}

	return data
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(reverseASlice(data))
}
