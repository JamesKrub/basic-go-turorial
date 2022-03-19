package main

import "fmt"

func addNewData(data map[string]int) {
	data["kate"] = 20
}
func main() {
	age := make(map[string]int)
	age["betty"] = 16

	fmt.Println(age)
	addNewData(age)
	fmt.Println(age)
}
