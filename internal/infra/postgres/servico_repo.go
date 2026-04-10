package postgres

import (
	"context"
	"database/sql"

	"github.com/maurilio-Deboin/erp-servico/internal/domain/servico"
)

// ServicoRepo é a implementação do repositório de serviços usando PostgreSQL.
type ServicoRepo struct {
	db *sql.DB
}

// NovoServicoRepo cria uma nova instância de ServicoRepo com a conexão ao banco de dados.
func NovoServicoRepo(db *sql.DB) *ServicoRepo {
	return &ServicoRepo{db: db}
}

// Save insere um novo serviço no banco de dados e retorna o ID gerado.
func (r *ServicoRepo) Save(ctx context.Context, s *servico.Servico) error {
	query := `
		INSERT INTO servicos (id, titulo, descricao, status, cliente_nome, criado_em, atualizado_em)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
		RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		s.Titulo, s.Descricao, s.Status, s.ClienteNome, s.CriadoEm, s.AtualizadoEm,
	).Scan(&s.ID)
}

func (r *ServicoRepo) FindByID(ctx context.Context, id string) (*servico.Servico, error) {
	query := `
		SELECT id, titulo, descricao, status, cliente_nome, criado_em, atualizado_em
		FROM servicos WHERE id = $1`

	s := &servico.Servico{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&s.ID, &s.Titulo, &s.Descricao, &s.Status,
		&s.ClienteNome, &s.CriadoEm, &s.AtualizadoEm,
	)
	if err == sql.ErrNoRows {
		return nil, servico.ErrNaoEncontrado
	}
	return s, err
}

func (r *ServicoRepo) FindAll(ctx context.Context) ([]*servico.Servico, error) {
	query := `
		SELECT id, titulo, descricao, status, cliente_nome, criado_em, atualizado_em
		FROM servicos ORDER BY criado_em DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var servicos []*servico.Servico
	for rows.Next() {
		s := &servico.Servico{}
		if err := rows.Scan(
			&s.ID, &s.Titulo, &s.Descricao, &s.Status,
			&s.ClienteNome, &s.CriadoEm, &s.AtualizadoEm,
		); err != nil {
			return nil, err
		}
		servicos = append(servicos, s)
	}
	return servicos, nil
}

func (r *ServicoRepo) Update(ctx context.Context, s *servico.Servico) error {
	query := `
		UPDATE servicos
		SET status = $1, atualizado_em = $2
		WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, s.Status, s.AtualizadoEm, s.ID)
	return err
}