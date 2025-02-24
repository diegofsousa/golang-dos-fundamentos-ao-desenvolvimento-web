package main

import "fmt"

func readMessages(ch <-chan string) {
	msg1 := <-ch
	msg2 := <-ch
	fmt.Println("Mensagens recebidas:", msg1, "e", msg2)
}

func main() {
	ch := make(chan string, 2)

	ch <- "Mensagem 1"
	ch <- "Mensagem 2"

	readMessages(ch)
}
