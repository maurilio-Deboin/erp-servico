package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// NovaConexao cria uma nova conexão com o banco de dados PostgreSQL
func NovaConexao(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("erro ao abrir conexão com banco:", err)
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}

	// Configura o pool de conexões
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Loga o sucesso da conexão
	log.Println("banco de dados conectado")
	return db
}