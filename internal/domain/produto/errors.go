package produto

import (
	"errors"
)

var (
	ErrNaoEncontrado 	= errors.New("Produto não encontrado")
	ErrNomeObrigatorio 	= errors.New("Nome é obrigatorio")
	ErrPrecoInvalido	= errors.New("Preço invalido")
)