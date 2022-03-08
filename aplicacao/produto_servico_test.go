package aplicacao_test

import (
	"testing"

	"github.com/adortolan/go-hexagonal/aplicacao"
	mock_aplicacao "github.com/adortolan/go-hexagonal/aplicacao/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProdutoServico_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	produto := mock_aplicacao.NewMockProdutoInterface(ctrl)
	persistencia := mock_aplicacao.NewMockProdutoPersistenciaInterface(ctrl)
	persistencia.EXPECT().Get(gomock.Any()).Return(produto, nil).AnyTimes()
	service := aplicacao.ProdutoServico{Persistence: persistencia}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, produto, result)
}

func TestProdutoServico_NovoProduto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	produto := mock_aplicacao.NewMockProdutoInterface(ctrl)
	persistencia := mock_aplicacao.NewMockProdutoPersistenciaInterface(ctrl)
	persistencia.EXPECT().Gravar(gomock.Any()).Return(produto, nil).AnyTimes()
	service := aplicacao.ProdutoServico{Persistence: persistencia}

	result, err := service.NovoProduto("Caneta")
	require.Nil(t, err)
	require.Equal(t, produto, result)
}
