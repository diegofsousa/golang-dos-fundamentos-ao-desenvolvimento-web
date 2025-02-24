package main

import (
	"fmt"
	"sync"
	"time"
)

func makeCoffe(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fazendo o café...")
	time.Sleep(2 * time.Second)
	fmt.Println("Café pronto!")
}

func toastBread(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fazendo a torrada...")
	time.Sleep(3 * time.Second)
	fmt.Println("Torrada pronta!")
}

func fryEggs(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fritando ovos...")
	time.Sleep(1 * time.Second)
	fmt.Println("Ovos prontos!")
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)

	// rodando de maneira concorrente
	go makeCoffe(&wg)
	go toastBread(&wg)
	go fryEggs(&wg)

	wg.Wait()

	fmt.Println("Café da manhã está pronto!")
	fmt.Printf("Tempo total: %v\n", time.Since(start))
}
