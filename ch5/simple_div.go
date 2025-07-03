package main

import "fmt"

func main() {
	fmt.Println(div(11, 23))
}

func div(num, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}
