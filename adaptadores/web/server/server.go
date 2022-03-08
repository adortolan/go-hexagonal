package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adortolan/go-hexagonal/adaptadores/web/handler"
	"github.com/adortolan/go-hexagonal/aplicacao"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service aplicacao.ProdutoPersistenciaInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProdutoHandlers(r, n, w.Service)
	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
