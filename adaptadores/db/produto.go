package db

import (
	"database/sql"

	"github.com/adortolan/go-hexagonal/aplicacao"
	_ "github.com/mattn/go-sqlite3"
)

type ProdutoDb struct {
	db *sql.DB
}

func NovoProdutoDb(db *sql.DB) *ProdutoDb {
	return &ProdutoDb{db: db}
}

func (p *ProdutoDb) Get(pId string) (aplicacao.ProdutoInterface, error) {
	var Produto aplicacao.Produto
	stmt, err := p.db.Prepare("select codigo, nome from produto where codigo=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(pId).Scan(&Produto.CODIGO, &Produto.NOME)
	if err != nil {
		return nil, err
	}
	return &Produto, nil
}

func (p *ProdutoDb) Gravar(produto aplicacao.ProdutoInterface) (aplicacao.ProdutoInterface, error) {
	var rows int
	p.db.QueryRow("select codigo from produto where codigo=?", produto.GetCodigo()).Scan(&rows)
	if rows == 0 {
		_, err := p.NovoProduto(produto)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.Atualizar(produto)
		if err != nil {
			return nil, err
		}
	}
	return produto, nil
}

func (p *ProdutoDb) NovoProduto(produto aplicacao.ProdutoInterface) (aplicacao.ProdutoInterface, error) {
	stmt, err := p.db.Prepare(`insert into produto(codigo, nome) values(?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(produto.GetCodigo(), produto.GetNome())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (p *ProdutoDb) Atualizar(produto aplicacao.ProdutoInterface) (aplicacao.ProdutoInterface, error) {
	_, err := p.db.Exec("update produto set nome=? where codigo=?", produto.GetNome(), produto.GetCodigo())
	if err != nil {
		return nil, err
	}
	return produto, nil
}
