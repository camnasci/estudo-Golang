package main

import (
	"fmt"
)

func main() {

	//i := 0

	/* for i < 10 {
		i++
		fmt.Println("Incrementando")
		time.Sleep(time.Second)
	}
	fmt.Println(i) */

	// for j := 0; j < 50; j += 5 {
	// 	fmt.Println("Incrementando j ", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"Camila", "Ana", "Paula"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Camila",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
