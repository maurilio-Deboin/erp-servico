package produto

import (
	"errors"
	"time"
)


// UnidadeMedida representa a medida utilizado no produto
type UnidadeMedida string

const(
	Unidade UnidadeMedida = "unidade"
	Metro	UnidadeMedida = "Metro"
	Litros	UnidadeMedida = "Litros"
	Kilos	UnidadeMedida = "Kilos"
)

// Produto faz parte do sistema da Serviço
type Produto struct {
	ID			string
	Nome 		string
	Descricao	string
	Unidade		UnidadeMedida
	PrecoCompra	float64
	PrecoVenda	float64
	CriadoEM	time.Time
	AtualizadoEm time.Time
}

//cria um novo produto
func New(nome,descricao string,unidade UnidadeMedida,precoCompra, precoVenda float64, criacaoEM, atualizadoEm time.Time) (*Produto,error) {
	if nome == ""{
		return nil, errors.New("Nome é obrigatorio")
	}
	if precoVenda < 0 {
		return nil, errors.New("O preço não pode ser negativo")
	}
	return &Produto{
		Nome: 			nome,
		Descricao: 		descricao,
		Unidade: 		unidade,
		PrecoCompra: 	precoCompra,
		PrecoVenda: 	precoVenda,
		CriadoEM: 		time.Now(),
		AtualizadoEm: 	time.Now(),
	}, nil
}