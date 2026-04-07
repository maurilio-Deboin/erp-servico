package produto

import (
	"context"
)

type Repository interface{
	Save(ctx context.Context, p *Produto) error
	FindByID(ctx context.Context, id string) (*Produto, error)
	FindAll(ctx context.Context) ([]*Produto, error)
	Update(ctx context.Context, p *Produto) error
}
