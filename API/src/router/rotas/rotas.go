package rotas

import (
	"api/src/middlewares"
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
	rotas = append(rotas, rotasPublicacoes...)

	for _, rotas := range rotas {

		if rotas.RequerAutenticacao {
			r.HandleFunc(rotas.URI, middlewares.Logger(middlewares.Autenticar(rotas.Funcao))).Methods(rotas.Metodo)
		} else {
			r.HandleFunc(rotas.URI, middlewares.Logger(rotas.Funcao)).Methods(rotas.Metodo)
		}

	}

	return r
}
