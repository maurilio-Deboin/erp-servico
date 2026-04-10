package financeiro

import (
	"context"
	"time"
)

type Repository interface {
	Save(ctx context.Context, l *Lancamento) error
	FindByID(ctx context.Context, id string) (*Lancamento, error)
	FindAll(ctx context.Context) ([]*Lancamento, error)
	ConsultarSaldo(ctx context.Context, inicio, fim time.Time) (float64, error)
}