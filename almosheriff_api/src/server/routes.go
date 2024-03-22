package server

import (
	"fmt"
	"net/http"
	"strconv"

	"almosheriff.api/src/database"
	"almosheriff.api/src/schemas"
	"almosheriff.api/src/utils"
	"github.com/gorilla/mux"
)

// Pedidos
func PedidosIndexHandle(w http.ResponseWriter, r *http.Request) error{
	var pedidos []schemas.PedidoOut
	ctx, err := database.IndexPedido(&pedidos)
	if err != nil {
		return err
	}
	utils.WriteJson(w, pedidos)
	ctx = ctx
	return nil
}

func PedidosCreateHandle(w http.ResponseWriter, r *http.Request) error{
	pedido_item, err := schemas.CreatePedido(r)
	// utils.WriteJson(w, pedido_item)
	utils.WriteJson(w, "modeled")
	if err != nil {
		// utils.WriteJson(w, err.Error())
		fmt.Fprintf(w, err.Error())
		return err
	}
	err = database.AddPedido(*pedido_item)
	if err != nil {
		utils.WriteJson(w, err.Error())
		return nil
	}
	return nil
}

func PedidosGetHandle(w http.ResponseWriter, r *http.Request) error{
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//handle inv path param
	}
	pedido, err := database.GetPedido(id)
	if err != nil {
		utils.WriteJson(w, err.Error())
		return nil
	}
	utils.WriteJson(w, pedido)
	return nil
}
