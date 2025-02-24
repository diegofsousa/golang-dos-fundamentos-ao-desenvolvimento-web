package main

import (
	"fmt"

	"github.com/diegofsousa/golang-curso/14-polimorfismo/datasources"
)

func main() {
	// datasource := datasources.NewMongoDB("localhost", "teste-polimorfismo")
	datasource := datasources.NewPostgreSQL("localhost", "pg-polimorfismo", 5432)
	saveNewName(&datasource, "diego")
	saveNewName(&datasource, "maria")
	getNames(&datasource)
}

func saveNewName(d datasources.Datasource, name string) {
	d.Save(name)
}

func getNames(d datasources.Datasource) {
	fmt.Println(d.GetAll())
}
