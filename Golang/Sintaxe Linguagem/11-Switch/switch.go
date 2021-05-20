package main

import (
	"fmt"
)

func finalDeSemana(numero int) string {
	switch numero {
	case 1:
		return "Sábado"
	case 2:
		return "Domingo"
	default:
		return "Dia nao localizado"
	}
}

func finalDeSemana2(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Sábado"
	case numero == 2:
		dia = "Domingo"
	default:
		dia = "Dia nao localizado"
	}

	return dia
}

func main() {
	fmt.Println("Switch")

	dia := finalDeSemana(1)
	fmt.Println(dia)

	dia2 := finalDeSemana2(2)
	fmt.Println(dia2)

	// fallthrough = joga passo o proximo case sem validar a condicao
}
