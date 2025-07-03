package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, remainder, err := diveAndReminder(20, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func diveAndReminder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}
