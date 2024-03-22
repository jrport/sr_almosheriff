package server

import (
	"net/http"

	"almosheriff.api/src/utils"
	"github.com/gorilla/mux"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)error

func MakeHandleFunc(f HandleFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if err:=f(w, r); err!=nil{
			utils.WriteJson(w, err.Error())
		}
	}
}

type ServerConf struct{
	port string
}

func NewApi(port string) *ServerConf{
	return &ServerConf{port: port}
}

func (server *ServerConf)Run(){
	router := mux.NewRouter()

	router.HandleFunc("/", MakeHandleFunc(PedidosIndexHandle))
	router.HandleFunc("/pedidos", MakeHandleFunc(PedidosIndexHandle))
	router.HandleFunc("/pedidos/{id}", MakeHandleFunc(PedidosGetHandle))
	router.HandleFunc("/pedidos/new", MakeHandleFunc(PedidosCreateHandle)).Methods("POST")

	http.ListenAndServe(server.port, router)
}
