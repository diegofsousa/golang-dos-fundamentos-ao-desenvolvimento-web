package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // declaração do channel

	go func() {
		fmt.Println("Escrevendo no channel...")
		ch <- "Mensagem 1" // escrevendo no channel
		fmt.Println("Mensagem enviada!")
	}()

	time.Sleep(2 * time.Second)

	go func() {
		msg := <-ch // lendo a mensagem do channel
		fmt.Printf("Recebido: %s\n", msg)
	}()

	time.Sleep(2 * time.Second)
}
