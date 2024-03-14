package schemas

import (
	_"time"

	"gorm.io/gorm"
)

type Model interface{
	GetType()string
}

type Pedido struct{
	gorm.Model
	ID uint
	Item string
	Quantidade int
	Atendido bool
	Compra int
	Origem int
}

func (p Pedido)GetType()string{
	return "Pedido"
}
