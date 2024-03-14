package database

import (
	"gorm.io/gorm"
	"jrport.almosheriff/schemas"
)

func SetPedidosTable(db *gorm.DB){
	db.AutoMigrate(&schemas.Pedido{})	
}
