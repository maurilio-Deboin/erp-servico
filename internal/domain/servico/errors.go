package servico

import "errors"

var (
	ErrNaoEncontrado    = errors.New("serviço não encontrado")
	ErrTituloObrigatorio = errors.New("título é obrigatório")
	ErrTransicaoInvalida = errors.New("transição de status inválida")
)