CREATE TABLE movimentacoes_estoque (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    produto_id  UUID NOT NULL REFERENCES produtos(id),
    tipo        TEXT NOT NULL CHECK (tipo IN ('entrada', 'saida')),
    quantidade  NUMERIC(10,3) NOT NULL,
    observacao  TEXT,
    criado_em   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);