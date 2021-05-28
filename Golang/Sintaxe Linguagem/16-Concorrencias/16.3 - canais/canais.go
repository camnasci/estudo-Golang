package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	go escrever("Ola Mundo!", canal)

	fmt.Println("Depois da funcao escrever comecar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	// fmt.Println("Fim do Programa")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)

	}

	close(canal)
}
