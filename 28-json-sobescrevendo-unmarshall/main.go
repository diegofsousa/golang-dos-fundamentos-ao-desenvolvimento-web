package main

import (
	"encoding/json"
	"fmt"
)

type Valor struct {
	Numero int
}

func (v Valor) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("Valor: %d", v.Numero))
}

func (v *Valor) UnmarshalJSON(data []byte) error {
	var raw string
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	_, err = fmt.Sscanf(raw, "Valor: %d", &v.Numero)
	return err
}

func main() {
	v := Valor{Numero: 42}
	jsonData, _ := json.Marshal(v)
	fmt.Println(string(jsonData))

	jsonInput := `"Valor: 53"`
	var vDecodificado Valor
	json.Unmarshal([]byte(jsonInput), &vDecodificado)
	fmt.Println("Valor decodificado:", vDecodificado.Numero)

}
