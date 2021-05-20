package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map")

	usuario := map[string]string{
		"nome":      "Camila",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[int]string{
		"nome": {
			1: "Camila",
			2: "Ana",
		},
		"sobrenome": {
			1: "Santos",
			2: "Silva",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["teste"] = map[int]string{
		3: "Souza",
	}
	fmt.Println(usuario2)

}
