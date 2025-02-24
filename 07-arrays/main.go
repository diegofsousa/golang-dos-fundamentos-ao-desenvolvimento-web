package main

import "fmt"

func main() {

	// forma básica
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 3
	numbers[2] = 67
	fmt.Println(numbers)

	// inicialização durante a declaração
	days := [3]string{"segunda", "terça", "quarta"}
	fmt.Println(days)

	// inferencia de valores
	notas := [...]int{55, 23, 67}
	fmt.Println(notas)
}
