package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randomsInt []int

	for i := 0; i <= 100; i++ {
		number := rand.Intn(100)
		randomsInt = append(randomsInt, number)
		for _, v := range randomsInt {
			fmt.Println(v)
		}
	}

}
