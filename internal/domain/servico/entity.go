package servico

import (
	"errors"
	"time"
)

// Status representa em que etapa o serviço está
type Status string

const (
	StatusAberto     Status = "aberto"
	StatusEmExecucao Status = "em_execucao"
	StatusEncerrado  Status = "encerrado"
)

// Servico é a entidade central do sistema
type Servico struct {
	ID          string
	Titulo      string
	Descricao   string
	Status      Status
	ClienteNome string // simples sem módulo CRM ainda
	CriadoEm   time.Time
	AtualizadoEm time.Time
}

// New cria um novo serviço já com valores padrão
func New(titulo, descricao, clienteNome string) (*Servico, error) {
	if titulo == "" {
		return nil, errors.New("título é obrigatório")
	}
	return &Servico{
		Titulo:      titulo,
		Descricao:   descricao,
		ClienteNome: clienteNome,
		Status:      StatusAberto,
		CriadoEm:   time.Now(),
		AtualizadoEm: time.Now(),
	}, nil
}

// Iniciar transiciona o serviço para em execução
func (s *Servico) Iniciar() error {
	if s.Status != StatusAberto {
		return errors.New("apenas serviços abertos podem ser iniciados")
	}
	s.Status = StatusEmExecucao
	s.AtualizadoEm = time.Now()
	return nil
}

// Encerrar finaliza o serviço
func (s *Servico) Encerrar() error {
	if s.Status != StatusEmExecucao {
		return errors.New("apenas serviços em execução podem ser encerrados")
	}
	s.Status = StatusEncerrado
	s.AtualizadoEm = time.Now()
	return nil
}