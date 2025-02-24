package main

import (
	"fmt"
)

func main() {

	// Loop for - exemplo básico
	fmt.Println("For básico")
	for i := 0; i < 5; i++ {
		fmt.Println("Valor de i: ", i)
	}

	// For como while
	fmt.Println("\n For como while")
	i := 0
	for i < 5 {
		fmt.Println("Valor de i: ", i)
		i++
	}

	// Loop infinito
	fmt.Println("\n Loop infinito")
	for {
		fmt.Println("loop infinito...")
		break
	}

	// range
	numbers := []int{10, 20, 30}

	fmt.Println("\n Usando o range para slices")
	for index, value := range numbers {
		fmt.Printf("indice: %d, valor: %d\n", index, value)
	}

	fmt.Println("\n Usando o range para maps")
	capitais := map[string]string{
		"Brasil": "Brasília",
		"França": "Paris",
		"Japão":  "Tókio",
	}

	for _, capital := range capitais {
		fmt.Printf("capital: %s\n", capital)
	}

}
