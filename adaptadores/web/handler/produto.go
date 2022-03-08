package handler

import (
	"net/http"

	"encoding/json"

	"github.com/adortolan/go-hexagonal/aplicacao"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProdutoHandlers(r *mux.Router, n *negroni.Negroni, service aplicacao.ProdutoServicoInterface) {
	r.Handle("/produto/{codigo}", n.With(
		negroni.Wrap(getProduto(service)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/produto", n.With(
		negroni.Wrap(NovoProduto(service)),
	)).Methods("POST", "OPTIONS")
}

func getProduto(service aplicacao.ProdutoServicoInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		codigo := vars["codigo"]
		produto, err := service.Get(codigo)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func NovoProduto(service aplicacao.ProdutoServicoInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var TProduto aplicacao.Produto
		err := json.NewDecoder(r.Body).Decode(&TProduto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		produto, err := service.NovoProduto(TProduto.NOME)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
