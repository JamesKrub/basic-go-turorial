package main

import "fmt"

func sum(nums ...int) int {
	var r int
	for _, num := range nums {
		r += num
	}
	return r
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
