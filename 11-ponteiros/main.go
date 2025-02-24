package main

import "fmt"

func main() {
	// & obtém o endereço de uma variável
	// * acessa o valor armazenado no endereço que o ponteiro aponta

	name := "diego"
	var pointer *string

	fmt.Println(name)
	fmt.Println(pointer)

	fmt.Println("\nendereços de memória das variáveis")
	fmt.Println(&name)
	fmt.Println(&pointer)

	fmt.Println("\npointer referenciando name")
	pointer = &name
	fmt.Println(pointer)
	fmt.Println(*pointer)

	fmt.Println("\nalteração do valor de name pelo ponteiro")
	*pointer = "diego fernando"
	fmt.Println(name)

	fmt.Println("\nalteração do valor de name pela propria variavel name")
	name = "maria"
	fmt.Println(*pointer)
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(pointer)
	fmt.Println(&pointer)

}
