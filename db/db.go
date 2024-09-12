package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func ConectaBanco() *sql.DB {

	once.Do(func() {
		conexao := "user=postgres password=password dbname=teste sslmode=disable host=localhost port=5432"
		db, err = sql.Open("postgres", conexao)

		if err != nil {
			log.Fatal("Erro ao conectar banco:", err)
			return
		}

		if err = db.Ping(); err != nil {
			log.Fatal("Erro ao verificar a conexão com o banco:", err)
			return
		}
	})

	return db
}
