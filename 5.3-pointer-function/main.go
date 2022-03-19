package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func assignValueToX(v *Vertex) {
	v.X = 10
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	assignValueToX(&v)
	fmt.Println(v)
}
