package aplicacao_test

import (
	"testing"

	"github.com/adortolan/go-hexagonal/aplicacao"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduto_EValido(t *testing.T) {
	produto := aplicacao.Produto{}
	produto.CODIGO = uuid.NewV4().String()
	produto.NOME = "Caneta"

	_, err := produto.EValido()
	require.Nil(t, err)
}
