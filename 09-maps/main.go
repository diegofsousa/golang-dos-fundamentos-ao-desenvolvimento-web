package main

import "fmt"

func main() {
	// forma basica de declaração
	var m map[string]int
	m = make(map[string]int)
	m["diego"] = 28
	m["maria"] = 27

	// recuperando um valor no map
	fmt.Println(m["diego"])
	fmt.Println(m["asdaisu"])

	age, exists := m["qweiu"]
	if exists {
		fmt.Println("diego existe no map")
		fmt.Println(age)
	}

	delete(m, "diego")
	fmt.Println(m)

	// declaração e inicialização juntos
	capitais := make(map[string]string)
	capitais["sao paulo"] = "sao paulo"
	capitais["ceara"] = "fortaleza"
	capitais["pernambuco"] = "recife"

	fmt.Println(capitais["pernambuco"])

	for estado, capital := range capitais {
		fmt.Println()
		fmt.Println(estado)
		fmt.Println(capital)
	}

	// declaração e inicialição com valores
	idade := map[string]int{"diego": 28, "maria": 27}
	fmt.Println(idade)

	// maps aninhados
	innerMaps := make(map[string]map[string]int)
	innerMaps["outer"] = make(map[string]int)
	innerMaps["outer"]["inner"] = 30
	fmt.Println(innerMaps)

}
