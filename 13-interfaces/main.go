package main

import "fmt"

type Mamifero interface {
	saudacao()
	hobbie()
}

type Humano struct {
	Nome      string
	Profissao string
}

func (h Humano) saudacao() {
	fmt.Println("Oi, meu nome é " + h.Nome + " e eu sou " + h.Profissao)
}

func (h Humano) hobbie() {
	fmt.Println("video game")
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) saudacao() {
	fmt.Println(c.Nome + "diz: au au")
}

func (c Cachorro) hobbie() {
	fmt.Println("brincar")
}

type Gato struct {
	Nome        string
	TempoDeSono int
}

func (g Gato) saudacao() {
	fmt.Println(g.Nome + "diz: miau")
}

func (g Gato) hobbie() {
	fmt.Printf("durmo umas %d horas", g.TempoDeSono)
}

func main() {
	diego := Humano{
		Nome:      "diego",
		Profissao: "engenheiro de software",
	}

	thor := Cachorro{
		Nome: "thor",
	}

	mingau := Gato{
		Nome:        "mingau",
		TempoDeSono: 12,
	}

	apresentacao(diego)
	apresentacao(thor)
	apresentacao(mingau)
}

func apresentacao(m Mamifero) {
	m.saudacao()
	m.hobbie()
}
