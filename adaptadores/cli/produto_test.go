package cli_test

import (
	"fmt"

	"testing"

	"github.com/adortolan/go-hexagonal/adaptadores/cli"
	mock_aplicacao "github.com/adortolan/go-hexagonal/aplicacao/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ProdutoNome := "Telefone Celular"
	ProdutoCodigo := "abc"

	produtoMock := mock_aplicacao.NewMockProdutoInterface(ctrl)
	produtoMock.EXPECT().GetCodigo().Return(ProdutoCodigo).AnyTimes()
	produtoMock.EXPECT().GetNome().Return(ProdutoNome).AnyTimes()

	service := mock_aplicacao.NewMockProdutoServicoInterface(ctrl)
	service.EXPECT().NovoProduto(ProdutoNome).Return(produtoMock, nil).AnyTimes()
	service.EXPECT().Get(ProdutoCodigo).Return(produtoMock, nil).AnyTimes()

	resultadoEsperado := fmt.Sprintf("Produto código %s com o nome %s foi criado com sucesso", ProdutoCodigo, ProdutoNome)
	result, err := cli.Run(service, "novoproduto", "", ProdutoNome)
	require.Nil(t, err)
	require.Equal(t, resultadoEsperado, result)

	resultadoEsperado = fmt.Sprintf("Produto codigo: %s\nNome: %s", ProdutoCodigo, ProdutoNome)
	result, err = cli.Run(service, "get", ProdutoCodigo, "")
	require.Nil(t, err)
	require.Equal(t, resultadoEsperado, result)
}
