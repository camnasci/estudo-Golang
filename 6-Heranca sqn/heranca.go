package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"joao", "paulo", 20, 160}

	es1 := estudante{p1, "Ciencia da Computacao", "FIAP"}
	fmt.Println(es1)
}
