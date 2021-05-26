package main

import (
	"fmt"
	"time"
)

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	go escrever("Ola Mundo!") //goroutines = nao espera terminar e executa o proximo comando
	escrever("Programando em GO!")
}
