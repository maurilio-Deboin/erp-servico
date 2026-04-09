package estoque

import "errors"

var (
	ErrSaldoInsuficiente  = errors.New("saldo insuficiente em estoque")
	ErrProdutoObrigatorio = errors.New("produto é obrigatório")
	ErrQuantidadeInvalida = errors.New("quantidade deve ser maior que zero")
)