package schemas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PedidoIn struct{
	Item string
	Quantidade int
	Origem int
}

type PedidoOut struct{
	Item string
	Quantidade int
	Origem int
	CreatedAt string
}

func NewPedidoOut(p *Pedido) *PedidoOut{
	return &PedidoOut{
		Item: p.Item,
		Quantidade: p.Quantidade,
		Origem: p.Origem,
		CreatedAt: fmt.Sprintf("%d-%d-%d", p.CreatedAt.Day(), p.CreatedAt.Month(), p.CreatedAt.Year()),
	}
}

func NewPedidoIn(r *http.Request) (*PedidoIn, error){
	var pedido PedidoIn
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil{
		return nil, err
	}
	return &pedido, nil
}

func CreatePedido(r *http.Request) (*Pedido, error){
	pedido_in, err := NewPedidoIn(r)	
	if err != nil {
		return nil, err
	}
	pedido_model := NewPedidoModel(*pedido_in)
	return pedido_model, nil
}
