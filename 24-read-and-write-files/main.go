package main

import (
	"bufio"
	"fmt"
	"os"
)

func simpleReadFile() {
	data, err := os.ReadFile("nomes.txt")
	if err != nil {
		fmt.Println("erro ao ler o arquivo:", err)
		return
	}
	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(data))
}

func lineReadFile() {
	file, err := os.Open("nomes.txt")
	if err != nil {
		fmt.Println("erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Linhas do meu arquivo:")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("erro durante a leitura:", err)
	}
}

func writeFile() {
	data := "Camaro\nBMW"
	err := os.WriteFile("carros.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("erro ao escrever no arquivo:", err)
		return
	}
	fmt.Println("Dados escritos no aquivo")
}

func writeAlreadyExistFile() {
	file, err := os.OpenFile("times.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("erro ao abrir o arquivo para escrita:", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString("\nFluminense\nInternacional\nFortaleza")
	if err != nil {
		fmt.Println("falha ao adicionar string no arquivo")
		return
	}
	fmt.Println("Dados adicionados com sucesso!")

}

func main() {
	fmt.Println("Leitura simples")
	simpleReadFile()
	fmt.Printf("\nLeitura linha a linha\n")
	lineReadFile()
	fmt.Printf("\nEscrevendo em arquivo\n")
	writeFile()
	fmt.Printf("\nEscrevendo no fim de um arquivo existente\n")
	writeAlreadyExistFile()
}
