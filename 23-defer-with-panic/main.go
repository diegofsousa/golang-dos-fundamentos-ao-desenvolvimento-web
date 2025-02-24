package main

import (
	"fmt"
	"os"
)

func writeFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	fmt.Fprintln(file, "Hello, world!")
	fmt.Println("Arquivo foi escrito corretamente.")

}

func main() {
	writeFile()
}
