package main

import (
	"fmt"
	"strings"
)

func main() {
	name := " Diego Fernando "
	fmt.Println("Original:", name)
	fmt.Println("Sem espaços:", strings.TrimSpace(name))
	fmt.Println("Em maiúsculas:", strings.ToUpper(name))
	fmt.Println("Em minúsculas:", strings.ToLower(name))
	fmt.Println("Substituição:", strings.ReplaceAll(name, "Fernando", "Lima"))

	name2 := "Maria"
	age := 28
	balance := 12998.56

	msg := fmt.Sprintf("\n\n\nNome: %s\nIdade: %d\nSaldo: %.2f", name2, age, balance)
	fmt.Println(msg)

}
