package main

import "fmt"

func main() {
loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto loop
		}
		fmt.Println(i)
	}
}
