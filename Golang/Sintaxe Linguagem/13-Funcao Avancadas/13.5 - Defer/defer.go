package main

import "fmt"

func calculaMedia(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado! ") // Adiar a execucao de algum código
	fmt.Println("Entrando na funcao ")

	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("Deu Bom!")
		return true
	}
	fmt.Println("Deu Ruim!")
	return false
}

func main() {
	fmt.Println(calculaMedia(7, 9))

}
