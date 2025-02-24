package main

import "fmt"

func sendMessage(ch chan<- string) {
	ch <- "Mensagem 1"
	ch <- "Mensagem 2"
}

func main() {
	ch := make(chan string, 2)

	go sendMessage(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
