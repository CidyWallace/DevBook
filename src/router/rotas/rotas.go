package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rotas := range rotas {
		r.HandleFunc(rotas.URI, rotas.Funcao).Methods(rotas.Metodo)
	}

	return r
}
