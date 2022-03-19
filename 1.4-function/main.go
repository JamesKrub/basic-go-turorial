package main

import (
	"fmt"
	"strings"
)

func title(name string) {
	msg := strings.ToUpper("Hello " + name)
	fmt.Println(msg)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	//title("Wat")
	//title("Nine")
	//num := add(10, 20)
	//
	//func(num int) {
	//	num = 50
	//	fmt.Println(num)
	//}(num)
	//
	//{
	//	num := 100
	//	fmt.Println(num)
	//}
	//
	//fmt.Println(num)
}
