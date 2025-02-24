package main

import "fmt"

func main() {
	fmt.Println("Operadores aritmeticos: ")
	sub := 80 - 6
	sum := 40 + 7
	div := 45 / 7
	mul := 56 * 23

	fmt.Println(sum)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mul)

	fmt.Println("\nOperadores relacionais: ")
	fmt.Println(34 > 4)

	age := 18

	fmt.Println(age > 18) // >, >=, ==, <, <=

	fmt.Println("\nOperadores de atribuição: ")
	x := 10
	x += 5 // x = x + 5
	x *= 20

	y := 20
	y -= 34
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("\nOperadores lógicos: ")

	userPass := "123456"
	isAdmin := true
	if userPass == "12345" || isAdmin == true { // && = e, || = ou
		fmt.Println("usuário logado")
	} else {
		fmt.Println("não permitido")
	}

	// Bit a Bit
	fmt.Println("\nOperadores bit a bit: ")
	a := 10 // 1010
	b := 3  // 0011
	result := a & b

	fmt.Println(result)

}
