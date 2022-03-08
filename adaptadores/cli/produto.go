package cli

import (
	"fmt"

	"github.com/adortolan/go-hexagonal/aplicacao"
)

func Run(service aplicacao.ProdutoServicoInterface, pAction string, pCodigo string, pNome string) (string, error) {

	var result = ""

	switch pAction {
	case "novoproduto":
		produto, err := service.NovoProduto(pNome)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Produto código %s com o nome %s foi criado com sucesso", produto.GetCodigo(), produto.GetNome())

	default:
		res, err := service.Get(pCodigo)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Produto codigo: %s\nNome: %s", res.GetCodigo(), res.GetNome())
	}
	return result, nil
}
