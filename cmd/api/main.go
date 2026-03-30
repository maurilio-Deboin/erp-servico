package main

import (
	"fmt"
	"github.com/maurilio-Deboin/erp-servico/configs"
	"github.com/maurilio-Deboin/erp-servico/internal/domain/servico"
)

func main() {
	cfg := configs.Load()

	// Só para testar se a entidade funciona
	s, err := servico.New("Instalação do piso", "Troca de bancada", "João Silva")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Serviço criado:", s.Titulo, "| Status:", s.Status)
	fmt.Println("ERP iniciando na porta", cfg.Port)
}