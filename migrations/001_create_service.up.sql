-- migrations/001_create_servicos.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE servicos (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    titulo        TEXT NOT NULL,
    descricao     TEXT,
    status        TEXT NOT NULL DEFAULT 'aberto',
    cliente_nome  TEXT,
    criado_em     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    atualizado_em TIMESTAMPTZ NOT NULL DEFAULT NOW()
);