package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"jrport.almosheriff/database"
	"jrport.almosheriff/schemas"
	"jrport.almosheriff/utils"
)

// Funcs de rotas para pedidos
func HandleIndexPedidos(w http.ResponseWriter, r *http.Request) error{
	pedidos, quantidade_pedido, err := database.IndexPedidos()
	if err != nil {
		return err
	}

	var pedidos_out []schemas.PedidoOut
	for i := 0; i<int(quantidade_pedido); i++ {
		pedidos_arr := *pedidos
		pedido_out := schemas.NewPedidoOut(&pedidos_arr[i])
		pedidos_out = append(pedidos_out, *pedido_out)	
	}

	utils.WriteJson(w, pedidos_out)
	return nil
}

func HandleGetPedido(w http.ResponseWriter, r *http.Request) error{
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid id for item")
		return err
	}
	pedido, err := database.GetPedido(id)
	if err != nil{
		return err
	}

	pedido_out := schemas.NewPedidoOut(pedido)
	utils.WriteJson(w, pedido_out)
	return nil
}

func HandleUpdatePedido(w http.ResponseWriter, r *http.Request) error{
	fmt.Fprintf(w, "bla")
	return nil
}

func HandleDeletePedido(w http.ResponseWriter, r *http.Request) error{
	fmt.Fprintf(w, "bla")
	return nil
}
