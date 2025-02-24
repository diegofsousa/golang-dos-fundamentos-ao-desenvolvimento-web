package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) showName() {
	fmt.Println(&p.name)
}

func main() {

	fmt.Println("ponteiros com funções")
	age := 25
	fmt.Println(&age)
	showAge(&age)
	fmt.Println(age)

	fmt.Println("\nponteiros com métodos")
	p := Person{
		name: "diego",
	}
	fmt.Println(&p.name)
	p.showName()

}

// passagem por valor = age int
// passagem por referência = age *int
func showAge(age *int) {
	fmt.Println(age)
	fmt.Println(*age)
	*age = 30
}
