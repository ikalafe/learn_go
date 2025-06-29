package main

import "fmt"

func main() {
	var b byte = 10
	var smallI int32 = 20
	var bigI uint64 = 30

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

}
