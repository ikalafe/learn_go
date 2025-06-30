package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("e", e)
}
