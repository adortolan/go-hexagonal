package aplicacao

import (
	"github.com/asaskevich/govalidator"

	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProdutoInterface interface {
	EValido() (bool, error)
	GetCodigo() string
	GetNome() string
}

type ProdutoServicoInterface interface {
	Get(codigo string) (ProdutoInterface, error)
	NovoProduto(nome string) (ProdutoInterface, error)
}

type ProdutoLeitor interface {
	Get(codigo string) (ProdutoInterface, error)
}

type ProdutoGrava interface {
	Gravar(Produto ProdutoInterface) (ProdutoInterface, error)
}

type ProdutoPersistenciaInterface interface {
	ProdutoLeitor
	ProdutoGrava
}

type Produto struct {
	CODIGO string `valid:"uuidv4"`
	NOME   string `valid:"required"`
}

func NovoProduto() *Produto {
	prod := Produto{
		CODIGO: uuid.NewV4().String(),
	}
	return &prod
}

func (p *Produto) GetCodigo() string {
	return p.CODIGO
}

func (p *Produto) GetNome() string {
	return p.NOME
}

func (p *Produto) EValido() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}
