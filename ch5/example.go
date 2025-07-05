package main

import "fmt"

func main() {
	x := 10
	faledUPdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}

// // *************************** ////
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}

	return total
}

//// *************************** ////

func faledUPdate(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}
