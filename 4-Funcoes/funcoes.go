package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	resultadoSoma := n1 + n2
	resultadoSub := n1 - n2

	return resultadoSoma, resultadoSub
}

func main() {
	soma := somar(10, 30)
	fmt.Println(soma)

	soma2, subtracao := calculosMatematicos(50, 20)
	fmt.Println(soma2, subtracao)
}
