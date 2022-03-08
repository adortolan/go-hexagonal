package main

import (
	"database/sql"

	"github.com/adortolan/go-hexagonal/aplicacao"

	db2 "github.com/adortolan/go-hexagonal/adaptadores/db"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	produtoDbAdapter := db2.NovoProdutoDb(db)
	produtoServico := aplicacao.ProdutoServico{Persistence: produtoDbAdapter}
	//produtoServico.NovoProduto("Borracha")
	produtoServico.Get("89edc6a3-1bc0-4987-82fa-f10f09d9b444")
	produtoServico.NovoProduto("Borracha Teste")

}
