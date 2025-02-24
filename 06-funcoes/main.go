package main

import (
	"fmt"
)

func main() {
	// função básica
	fmt.Println("função básica")
	r := sum(3, 4)
	fmt.Println(r)

	fmt.Println("\nretornando multiplos valores")
	_, r = divide(10, 3)
	fmt.Println(r)

	fmt.Println("\nfunção sem retorno")
	hello("diego")

	fmt.Println("\nfunção anonima")

	anon := func(a, b int) int {
		return a + b
	}

	fmt.Println(anon(6, 8))
	fmt.Println(anon(12, 8))

	fmt.Println("\nclosure")
	nextOdd := oddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	newNextOdd := oddGenerator()
	fmt.Println(newNextOdd())

	fmt.Println("\nfunção variadica")
	fmt.Println(max(1, 8, 5, 45, 2))

}

func hello(name string) {
	fmt.Println(name)
}

func sum(a, b int) int {
	return a + b
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func oddGenerator() func() int {
	i := -1
	return func() int {
		i += 2
		return i
	}
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	bigger := numbers[0]
	for _, num := range numbers {
		if num > bigger {
			bigger = num
		}
	}
	return bigger
}
