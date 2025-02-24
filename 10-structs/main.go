package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Email   string
	Address Address
	Employment
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Employment struct {
	Position string
}

func NewPerson(name, email, street, city, zip, position string, age int) Person {
	p := Person{
		Name:  name,
		Age:   age,
		Email: email,
		Address: Address{
			Street: street,
			City:   city,
			Zip:    zip,
		},
	}

	p.Position = position

	return p
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func main() {
	p := NewPerson("Diego", "diego@email.com", "Rua do Golang", "São Paulo", "252134-2309", "software engineer", 28)
	p1 := NewPerson("Maria", "maria@email.com", "Rua do Java", "rio de Janeiro", "252134-2309", "data engineer", 27)

	persons := []Person{}
	pensonsSlice := append(persons, p)
	pensonsSlice = append(pensonsSlice, p1)

	for _, p := range pensonsSlice {
		fmt.Println(p.Greet())
	}

	// diegoGreet := p.Greet()
	// fmt.Println(diegoGreet)
}
