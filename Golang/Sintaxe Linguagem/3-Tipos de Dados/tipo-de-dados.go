package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int64 = 1000000000000
	fmt.Println(numero)

	var numero2 uint = 1000000000000
	fmt.Println(numero2)

	//alias - int32 = rune
	var numero3 rune = 122323
	fmt.Println(numero3)

	//alias - int8 = rune
	var numero4 byte = 122
	fmt.Println(numero4)

	//float32, float64
	//sempre com ponto nas casas decimais

	//strings

	//char traz o valor da tabela ASC

	//bool (true e false)

	//error
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
