package main

import (
	"fmt"
)

func main() {
	var nome string = "Camila Silva"
	nome2 := "Ana"

	fmt.Println(nome)
	fmt.Println(nome2)

	var (
		nome3 string = "sjksjdsd"
		nome4 string = "nome4"
	)

	fmt.Println(nome3, nome4)

	nome5, nome6 := "Nome5", "nome6"
	fmt.Println(nome5, nome6)
}
