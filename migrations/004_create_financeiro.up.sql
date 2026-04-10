CREATE TABLE lancamentos (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    descricao   TEXT NOT NULL,
    tipo        TEXT NOT NULL CHECK (tipo IN ('receita', 'despesa')),
    valor       NUMERIC(10,2) NOT NULL,
    servico_id  UUID REFERENCES servicos(id) ON DELETE SET NULL,
    criado_em   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);