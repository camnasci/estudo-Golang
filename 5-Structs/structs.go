package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Struct")

	var u usuario
	u.nome = "Camila"
	u.idade = 27

	fmt.Println(u)

	enderecoExemplo := endereco{"Teste Endereco", 0}

	us2 := usuario{"Andre", 40, enderecoExemplo}
	fmt.Println(us2)

	us3 := usuario{nome: "Davi"}
	fmt.Println(us3)
}
