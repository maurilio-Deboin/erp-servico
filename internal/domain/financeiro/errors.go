package financeiro

import "errors"

var (
	ErrNaoEncontrado      = errors.New("lançamento não encontrado")
	ErrDescricaoObrigatorio = errors.New("descrição é obrigatória")
	ErrValorInvalido      = errors.New("valor deve ser maior que zero")
	ErrTipoInvalido       = errors.New("tipo inválido")
)