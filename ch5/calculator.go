package main

import (
	"fmt"
	"strconv"
)

func main() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"tow", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression: ", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operation: ", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }
