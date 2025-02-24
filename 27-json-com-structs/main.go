package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Employee struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Position string   `json:"position"`
	Skills   []string `json:"skills"`
	Adress   Address  `json:"address"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func validateEmployee(e Employee) error {
	if e.Name == "" {
		return errors.New("o nome do funcionário é obrigatório")
	}
	if e.Position == "" {
		return errors.New("o cargo do funcionário é obrigatório")
	}
	return nil
}

func main() {
	jsonData := `{
		"id":1,
		"name":"AAA",
		"position":"ASAS",
		"skills":["Go","Python"],
		"address": {
			"city":"Osasco",
			"country":"Brasil"
		}
	}`

	var employee Employee
	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println("Falha ao decodificar Json:", err)
		return
	}

	err = validateEmployee(employee)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Funcionario:")
	fmt.Println(employee)
}
