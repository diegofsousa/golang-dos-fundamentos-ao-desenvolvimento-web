package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchURL(wokerID int, url string) {
	fmt.Printf("Worker %d: começando uma requisição para %s\n", wokerID, url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Worker %d: erro ao acessar %s: %v\n", wokerID, url, err)
	}
	defer resp.Body.Close()
	fmt.Printf("Worker %d: concluído para %s com status %d\n", wokerID, url, resp.StatusCode)
}

func worker(id int, urls <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range urls {
		fetchURL(id, url)
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
	}

	const numWorkers = 3 // número de workers no pool
	urlChannel := make(chan string, len(urls))
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, urlChannel, &wg)
	}

	for _, url := range urls {
		urlChannel <- url
	}

	close(urlChannel)

	wg.Wait()
	fmt.Println("Todas as requisições foram concluídas!")

}
