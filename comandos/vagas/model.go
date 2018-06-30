package vagas

import "time"

type Vaga struct {
	ID          int
	Nome        string
	Cargo       string
	Descricao   string
	Empresa     string
	Salario     string
	Ativa       bool
	dataCriacao time.Time
}
