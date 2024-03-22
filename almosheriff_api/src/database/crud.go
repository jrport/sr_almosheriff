package database

import (
	"almosheriff.api/src/schemas"
	"gorm.io/gorm"
)

func AddPedido(pedido schemas.Pedido) error{
	session, err := GetSession()	
	if err != nil {
		return err
	}
	result := session.Create(&pedido)
	err = result.Error
	return err
}

func IndexPedido(pedidos_out *[]schemas.PedidoOut) (*gorm.DB, error){
	var pedidos []schemas.Pedido
	session, err := GetSession()	
	if err != nil {
		return nil, err
	}
	ctx := session.Find(&pedidos)
	for _, pedido := range pedidos {
		pedido_out := *schemas.NewPedidoOut(&pedido)	
		*pedidos_out = append(*pedidos_out, pedido_out)
	}
	return ctx, nil
}

func GetPedido(id int) (*schemas.PedidoOut, error){
	var pedido_out *schemas.PedidoOut
	session, err := GetSession()	
	if err != nil {
		return nil, err
	}
	pedido_obj := schemas.Pedido{}
	ctx := session.First(&pedido_obj, id)
	if ctx.Error != nil {
		return nil, ctx.Error
	}	
	pedido_out = schemas.NewPedidoOut(&pedido_obj)
	return pedido_out, nil
}
