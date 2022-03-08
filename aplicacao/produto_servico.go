package aplicacao

type ProdutoServico struct {
	Persistence ProdutoPersistenciaInterface
}

func NovoProdutoServico(persistencia ProdutoPersistenciaInterface) *ProdutoServico {
	return &ProdutoServico{Persistence: persistencia}
}

func (s *ProdutoServico) Get(codigo string) (ProdutoInterface, error) {
	produto, err := s.Persistence.Get(codigo)
	if err != nil {
		return nil, err
	}
	return produto, nil
}

func (s *ProdutoServico) NovoProduto(pNome string) (ProdutoInterface, error) {
	produto := NovoProduto()
	produto.NOME = pNome
	_, err := produto.EValido()
	if err != nil {
		return &Produto{}, err
	}
	result, err := s.Persistence.Gravar(produto)
	if err != nil {
		return &Produto{}, err
	}
	return result, nil
}
