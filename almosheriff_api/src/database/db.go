package database

import (
	"fmt"

	"almosheriff.api/src/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConn struct{
	dsn string
	config gorm.Config
}

func NewDbConn(port int, m ...int) *DbConn{
	dsn := fmt.Sprintf("host=localhost user=jrporto password=26072004 dbname=almosheriff port=%d sslmode=disable", port)
	if m != nil {
		return &DbConn{dsn: dsn, config: gorm.Config{SkipDefaultTransaction: true}}
	}
	return &DbConn{dsn: dsn, config: gorm.Config{}}
}

func GetSession() (*gorm.DB, error){
	dbCon := NewDbConn(5000)
	db, err := gorm.Open(postgres.Open(dbCon.dsn), &dbCon.config)
	if err != nil {
		return nil, err
	}
	println("Sucessfully connected!")
	return db, nil
}

func (dbConn DbConn)SetupTable() error{
	db, err := GetSession()
	if err != nil {
		return err
	}
	db.AutoMigrate(&schemas.Pedido{})
	return nil
}
