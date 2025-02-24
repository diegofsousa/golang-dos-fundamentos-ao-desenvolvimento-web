package main

import (
	"fmt"
	"sync"
)

var (
	counter int // recurso compartilhado
	mutex   sync.Mutex
)

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go incrementCounter(&wg)
	}

	wg.Wait()
	fmt.Println(counter)

}
