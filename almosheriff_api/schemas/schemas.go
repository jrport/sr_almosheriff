package schemas

type PedidoOut struct {
	Item string
	DataPedido string
	Quantidade int
	Origem int
}

func NewPedidoOut(pedido *Pedido) *PedidoOut{
	return &PedidoOut{
		Item: pedido.Item,
		Quantidade: pedido.Quantidade,
		Origem: pedido.Origem,
		DataPedido: pedido.CreatedAt.Format("January 02, 2006"),
	}
}
