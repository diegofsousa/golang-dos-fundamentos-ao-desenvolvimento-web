package main

import "fmt"

func main() {
	// if
	age := 33

	if age >= 18 {
		fmt.Println("você é maior de idade")
	}

	// if-else
	if age >= 18 {
		fmt.Println("você é maior de idade")
	} else {
		fmt.Println("você é menor de idade")
	}

	// if-else-if

	if age < 13 {
		fmt.Println("você é uma criança")
	} else if age < 18 {
		fmt.Println("você é um adolescente")
	} else {
		fmt.Println("você é um adulto")
	}

	day := "domingo"

	if day == "segunda" || day == "terça" || day == "quarta" || day == "quinta" || day == "sexta" {
		fmt.Println("hoje é um dia util")
	} else if day == "sábado" || day == "domingo" {
		fmt.Println("hoje é fim de semana")
	} else {
		fmt.Println("dia desconhecido")
	}

	// switch case
	switch day {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("hoje é um dia util")
	case "sábado", "domingo":
		fmt.Println("hoje é fim de semana")
	default:
		fmt.Println("dia desconhecido")
	}

}
