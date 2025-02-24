package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonData := `{"name":"Diego", "attributes": {"height":1.7, "skills":["Go", "Python"]}}`

	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {
		fmt.Println("falha ao decodificar o Json:", err)
		return
	}

	fmt.Println("Nome:", data["name"])
	attributes := data["attributes"].(map[string]interface{})
	fmt.Println("Altura:", attributes["height"])
	fmt.Println("Habilidades:", attributes["skills"])

}
