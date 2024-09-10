package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConectaBanco() {

	conexao := "user=postgres password=password dbname=teste sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Fatal("Erro ao conectar banco:", err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao verificar a conexão com o banco:", err)
		return
	}

	DB = db

}
