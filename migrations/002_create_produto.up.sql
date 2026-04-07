create extension if not exists "uuid-ossp";

create table produtos (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nome          TEXT NOT NULL,
    descricao     TEXT,
    unidade       TEXT,
    preco_compra  NUMERIC(10, 2) not null default 0,
    preco_venda   NUMERIC(10, 2) not null default 0,
    criado_em     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    atualizado_em TIMESTAMPTZ NOT NULL DEFAULT NOW()
)