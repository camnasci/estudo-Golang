package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo!")
		waitGroup.Done() //-1
	}()

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //esperar as goroutines executarem para finalizar

}
