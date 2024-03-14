package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	port string
}

func NewApiServer(addr string)*ApiServer{
	return &ApiServer{port: addr}
}

type ApiFunc func(w http.ResponseWriter, r *http.Request)error

func MakeHttpHandle(f ApiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err := f(w, r); err != nil{
			fmt.Fprintf(w, err.Error())
		}
	}
}

func (api ApiServer)Run(){
	router := mux.NewRouter()

	router.HandleFunc("/pedidos", MakeHttpHandle(HandleIndexPedidos))
	router.HandleFunc("/pedidos/{id}", MakeHttpHandle(HandleGetPedido))


	http.ListenAndServe(api.port, router)
}
