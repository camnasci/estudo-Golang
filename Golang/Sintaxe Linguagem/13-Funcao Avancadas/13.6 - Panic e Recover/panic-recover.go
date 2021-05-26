package main

import "fmt"

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso! ")
	}
}

func calculaMedia(n1, n2 float32) bool {
	defer recuperaExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media > 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(calculaMedia(6, 6))
	fmt.Println("Pós Execução!")
}
