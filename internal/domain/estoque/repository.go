package estoque

import "context"

type Repository interface {
	RegistrarMovimentacao(ctx context.Context, m *Movimentacao) error
	ConsultarSaldo(ctx context.Context, produtoID string) (float64, error)
	ListarMovimentacoes(ctx context.Context, produtoID string) ([]*Movimentacao, error)
}