package schemas

import "gorm.io/gorm"

type Pedido struct{
	gorm.Model
	Item string
	Origem int
	Quantidade int
}

func NewPedidoModel(pedido PedidoIn) *Pedido{
	return &Pedido{
		Item: pedido.Item,
		Quantidade: pedido.Quantidade,
		Origem: pedido.Origem,
	}
}
