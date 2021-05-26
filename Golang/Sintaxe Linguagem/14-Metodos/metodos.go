package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("metodo salvar")
}

func main() {
	us1 := usuario{"camila", 27}
	fmt.Println(us1)

	us1.salvar()
}
