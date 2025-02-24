package main

import "fmt"

const Name5 = "constante: diego"

func main() {
	// usando o var
	var name1 string
	name1 = "usando o var: diego"

	// iniciando valor com var
	var name2 string = "iniciando valor com o var: diego"

	// inferencia de tipo
	var name3 = "inferencia de tipo: diego"

	// declaração curta
	name4 := "declaração curta: diego"

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(Name5)

	var age int8 = 28
	var pi float32 = 3.14  // 6-7 digitos de precisão
	var pi2 float64 = 3.14 // 15-16 digitos de precisão

	fmt.Println(age)
	fmt.Println(pi)
	fmt.Println(pi2)

	var isEnabled bool = true
	fmt.Println(isEnabled)

	var a, b, c, d = "a", 50, true, 5.60

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
