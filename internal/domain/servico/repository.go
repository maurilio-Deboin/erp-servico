package servico

import "context"

// Repository define o contrato de persistência
// A implementação real fica em internal/infra/postgres/
type Repository interface {
	Save(ctx context.Context, s *Servico) error
	FindByID(ctx context.Context, id string) (*Servico, error)
	FindAll(ctx context.Context) ([]*Servico, error)
	Update(ctx context.Context, s *Servico) error
}