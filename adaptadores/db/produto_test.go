package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/adortolan/go-hexagonal/adaptadores/db"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	CreateTable(Db)
	CriarProduto(Db)
}

func CreateTable(db *sql.DB) {
	table := `CREATE TABLE PRODUTO("CODIGO" string, "NOME" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func CriarProduto(Db *sql.DB) {
	insert := `insert into produto values("abc", "Lapis")`
	stmt, err := Db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProdutoDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	produtoDb := db.NovoProdutoDb(Db)
	produto, err := produtoDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Lapis", produto.GetNome())
}
