package main

import "fmt"

func main() {
	printWhatYouSend("Hello, World")
	sumInt(1, 2)
	myValue(1.2)
}

func printWhatYouSend(i interface{}) {
	fmt.Println(i)
}

func sumInt(i, j interface{}) {
	num1 := i.(int)
	num2 := j.(int)
	fmt.Println(num1 + num2)
}

func myValue(myInterface interface{}) {
	switch v := myInterface.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		fmt.Printf("Integer: %v", v)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		fmt.Printf("Float64: %v", v)
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		fmt.Printf("String: %v", v)
	default:
		// And here I'm feeling dumb. ;)
		fmt.Printf("I don't know, ask stackoverflow.")
	}
}
