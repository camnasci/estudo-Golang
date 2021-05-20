package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de Controle")

	numero := 395

	if numero%2 == 0 {
		fmt.Println("É par")
	} else {
		fmt.Println("É impar")
	}

	//If Infinity
	numero2 := 300

	if outroNumero := numero2; outroNumero%2 == 0 {
		fmt.Println("É par")
	} else {
		fmt.Println("É impar")
	}

}
