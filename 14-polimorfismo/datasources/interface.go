package datasources

type Datasource interface {
	Save(name string)
	GetAll() []string
}
