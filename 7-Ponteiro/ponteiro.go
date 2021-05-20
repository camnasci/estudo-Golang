package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var p1 int
	var p2 *int

	p1 = 100
	p2 = &p1

	fmt.Println(p1, p2)
	fmt.Println(p1, *p2) //desreferenciacao
}
