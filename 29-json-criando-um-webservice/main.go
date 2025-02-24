package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requisição foi recebida!")
	var person Person
	// função para decodificar o json
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Erro ao processar Json", http.StatusBadRequest)
		return
	}
	fmt.Println(person)
	fmt.Fprintf(w, "Recebido: %+v", person)
}

func main() {
	fmt.Println("Servidor escutando na porta 8080")
	http.HandleFunc("/submit", handler)
	http.ListenAndServe(":8080", nil)

}
