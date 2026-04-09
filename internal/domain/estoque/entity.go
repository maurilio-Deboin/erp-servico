package estoque

import (
	"errors"
	"time"
)

type TipoMovimentacao string

const (
	Entrada TipoMovimentacao = "Entrada"
	Saida   TipoMovimentacao = "Saida"
)

// Movimentacao representa uma movimentacao de estoque
type Movimentacao struct {
	ID			string
	ProdutoID	string
	Tipo		TipoMovimentacao
	Quantidade	int
	Observacao	string
	CriadoEM	time.Time
	AtualizadoEm time.Time
}

// NovaMovimentacao cria uma nova movimentacao de estoque
func NovaMovimentacao(produtoID string, tipo TipoMovimentacao, quantidade int, observacao string) (*Movimentacao, error) {
	if produtoID == "" {
		return nil, errors.New("Produto é obrigatório")
	}
	if quantidade <= 0 {
		return nil, errors.New("Quantidade deve ser maior que zero")
	}
	return &Movimentacao{
		ProdutoID:  produtoID,
		Tipo:       tipo,
		Quantidade:  quantidade,
		Observacao: observacao,
		CriadoEM:   time.Now(),
		AtualizadoEm: time.Now(),
	}, nil
}	

