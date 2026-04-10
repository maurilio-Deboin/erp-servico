package postgres

import (
	"context"
	"database/sql"

	"github.com/maurilio-Deboin/erp-servico/internal/domain/produto"
)

// ProdutoRepo é a implementação do repositório de produtos usando PostgreSQL.
// Ele fornece métodos para salvar, buscar e atualizar produtos no banco de dados.
// Obs: irei aplicar o padrão Repository em todos os repositórios para manter a consistência e facilitar a manutenção do código.
type ProdutoRepo struct {
	db *sql.DB
}

// NovoProdutoRepo cria uma nova instância de ProdutoRepo com a conexão ao banco de dados.
func NovoProdutoRepo(db *sql.DB) *ProdutoRepo {
	return &ProdutoRepo{db: db}
}

// Save insere um novo produto no banco de dados e retorna o ID gerado.
func (r *ProdutoRepo) Save(ctx context.Context, p *produto.Produto) error {
	query := `
		INSERT INTO produtos (id, nome, descricao, unidade, preco_compra, preco_venda, criado_em, atualizado_em)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	return r.db.QueryRowContext(ctx, query,
		p.Nome, p.Descricao, p.Unidade,
		p.PrecoCompra, p.PrecoVenda,
		p.CriadoEM, p.AtualizadoEm,
	).Scan(&p.ID)
}
// FindByID busca um produto pelo seu ID no banco de dados. Retorna o produto encontrado ou um erro caso não exista.
func (r *ProdutoRepo) FindByID(ctx context.Context, id string) (*produto.Produto, error) {
	query := `
		SELECT id, nome, descricao, unidade, preco_compra, preco_venda, criado_em, atualizado_em
		FROM produtos WHERE id = $1`

	p := &produto.Produto{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.Nome, &p.Descricao, &p.Unidade,
		&p.PrecoCompra, &p.PrecoVenda,
		&p.CriadoEM, &p.AtualizadoEm,
	)
	if err == sql.ErrNoRows {
		return nil, produto.ErrNaoEncontrado
	}
	return p, err
}

// FindAll retorna uma lista de todos os produtos cadastrados no banco de dados, ordenados por nome.
func (r *ProdutoRepo) FindAll(ctx context.Context) ([]*produto.Produto, error) {
	query := `
		SELECT id, nome, descricao, unidade, preco_compra, preco_venda, criado_em, atualizado_em
		FROM produtos ORDER BY nome`
	// Executar a consulta e obter as linhas resultantes
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var produtos []*produto.Produto
	for rows.Next() {
		p := &produto.Produto{}
		if err := rows.Scan(
			&p.ID, &p.Nome, &p.Descricao, &p.Unidade,
			&p.PrecoCompra, &p.PrecoVenda,
			&p.CriadoEM, &p.AtualizadoEm,
		); err != nil {
			
			return nil, err
		}
		produtos = append(produtos, p)
	}
	return produtos, nil
}

// Update atualiza os dados de um produto existente no banco de dados.

func (r *ProdutoRepo) Update(ctx context.Context, p *produto.Produto) error {
	query := `
		UPDATE produtos
		SET nome = $1, descricao = $2, preco_compra = $3, preco_venda = $4, atualizado_em = $5
		WHERE id = $6`
	_, err := r.db.ExecContext(ctx, query,
		p.Nome, p.Descricao, p.PrecoCompra, p.PrecoVenda, p.AtualizadoEm, p.ID,
	)
	return err
}