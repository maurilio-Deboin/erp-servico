package produto

import (
	"errors"
)

var (
	errNãoEncontrado 	= errors.New("Produto não encontrado")
	errNomeObrigatorio 	= errors.New("Nome é obrigatorio")
	errPrecoInvalido	= errors.New("Preço invalido")
)