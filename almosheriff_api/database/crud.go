package database

import (
	"jrport.almosheriff/schemas"
)

func IndexPedidos() (*[]schemas.Pedido, int, error){
	dbConn, err := GetConnection()
	if err != nil{
		return nil, 0, err
	}
	var itens []schemas.Pedido
	ctx := dbConn.Find(&itens)	
	quant := int(ctx.RowsAffected)
	return &itens, quant, err
}

func GetPedido(id int) (*schemas.Pedido, error){
	dbConn, err := GetConnection()
	if err != nil {
		return nil, err
	}
	item := &schemas.Pedido{ID: uint(id)}
	dbConn.Find(item)
	return item, nil
}

func DeletePedido(id int) (int, error){
	dbConn, err := GetConnection()
	if err != nil {
		return 0, err
	}
	item := &schemas.Pedido{ID: uint(id)}
	ctx := dbConn.Delete(item)
	quant := int(ctx.RowsAffected)
	return quant, nil
}

