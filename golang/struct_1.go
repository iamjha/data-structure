package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v1 := new(Vertex)
	v1.X = 4
	v1.Y = 3
	v.X = 8
	v2 := &v
	v2.X = 5
	fmt.Println(v)
	fmt.Println(v1)
}
