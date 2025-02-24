package main

import "fmt"

func main() {

	// declaração básica de slices
	var numbers []int
	fmt.Println(numbers)

	// inicialização de slices
	days := []string{"segunda", "terça"}
	fmt.Println(days)

	// adicionando um novo valor ao slice
	numbers2 := []int{78, 90, 23}
	fmt.Println(numbers2)
	numbers2 = append(numbers2, 60, 90, 34, 21)
	fmt.Println(numbers2)
	numbers2 = append(numbers2, 123)
	fmt.Println(numbers2)

	// subslices
	numbers3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numbers3Subslice := numbers3[3:6]
	fmt.Println(numbers3Subslice)

	numbers3Sub1 := numbers3[:2]
	fmt.Println(numbers3Sub1)
	numbers3Sub2 := numbers3[3:10]
	fmt.Println(numbers3Sub2)

	// deletando uma posição através de subslices
	numbersDelete3Index := append(numbers3[:2], numbers3[3:]...)
	fmt.Println(numbersDelete3Index)

	// slices multidimensionais
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matriz)

	numbers4 := make([]int, 5, 10)
	fmt.Println(numbers4)
	numbers4[0] = 9
	fmt.Println(numbers4)
	fmt.Println(len(numbers4))
	fmt.Println(cap(numbers4))

	numbers5 := make([]int, 5, 10)
	fmt.Println(cap(numbers5))
	for i := 0; i < 6; i++ {
		numbers5 = append(numbers5, i)
	}

	fmt.Println(numbers5)
	fmt.Println(len(numbers5))
	fmt.Println(cap(numbers5))

}
