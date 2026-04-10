package financeiro

import(
	"errors"
	"time"
)


type TipoLancamento string

const (
	Receita TipoLancamento = "receita"
	Despesa TipoLancamento = "despesa"
)



type Lancamento struct {
	ID          string
	Descricao   string
	Valor       float64
	Tipo        TipoLancamento
	Data        time.Time
	ServicoID  string
	CriadoEm	time.Time

}

func NovoLancamento(descricao string, tipo TipoLancamento, valor float64, servicoID string) (*Lancamento, error) {
	if descricao == "" {
		return nil, errors.New("descrição é obrigatória")
	}
	if valor <= 0 {
		return nil, errors.New("valor deve ser maior que zero")
	}
	if tipo != Receita && tipo != Despesa {
		return nil, errors.New("tipo inválido, use receita ou despesa")
	}
	return &Lancamento{
		Descricao: descricao,
		Tipo:      tipo,
		Valor:     valor,
		ServicoID: servicoID,
		CriadoEm:  time.Now(),
	}, nil
}