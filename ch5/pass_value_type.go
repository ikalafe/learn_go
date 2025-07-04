package main

import "fmt"

func main() {
	p := person{}
	i := 2
	s := "Goodbye"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}
