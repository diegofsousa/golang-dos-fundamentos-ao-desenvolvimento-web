package datasources

import (
	"fmt"
)

var namesPgDB []string

type PostgreSQL struct {
	host     string
	port     int
	database string
}

func NewPostgreSQL(host, database string, port int) PostgreSQL {
	return PostgreSQL{
		host:     host,
		database: database,
		port:     port,
	}
}

func (p *PostgreSQL) Save(name string) {
	namesPgDB = append(namesPgDB, name)
	fmt.Println("dado persistido no postgres")
}

func (p PostgreSQL) GetAll() []string {
	return namesPgDB
}
