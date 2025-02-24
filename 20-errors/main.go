package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não é possivel dividir por 0")
	}
	return a / b, nil
}

func openFile(name string) error {
	if name == "" {
		return fmt.Errorf("erro ao abrir arquivo: %w", errors.New("nome do arquivo está vazio"))
	}
	return nil
}

func main() {
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(result)

	fmt.Println("adicionando contexto aos erros")

	err = openFile("")
	if err != nil {
		fmt.Println(err)
	}

}
