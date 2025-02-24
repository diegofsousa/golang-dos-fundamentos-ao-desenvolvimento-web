package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 3)

	go func() {
		fmt.Println("Escrevendo no channel com buffer")
		ch <- "Mensagem 1"
		fmt.Println("Mensagem 1 enviada!")
		ch <- "Mensagem 2"
		fmt.Println("Mensagem 2 enviada!")
		ch <- "Mensagem 3"
		fmt.Println("Mensagem 3 enviada!")
		ch <- "Mensagem 4"
		fmt.Println("Mensagem 4 enviada!")
	}()

	time.Sleep(2 * time.Second)

	go func() {
		for i := 0; i < 4; i++ {
			msg := <-ch
			fmt.Printf("Recebido: %s\n", msg)
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
}
