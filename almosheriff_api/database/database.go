package database

import (
	_"database/sql"
	_"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct{
	dsn string
}

type ConnectionConfig struct{
	Manual bool
}

const URI string = "postgres://jrporto:26072004@localhost:5000/inventario"

func NewPostgresSession(link string) *DbConnection{
	return &DbConnection{dsn: link}
}

func (dbConn *DbConnection)Connect(config *ConnectionConfig) (*gorm.DB, error){
	if config.Manual == true{
		db, err := gorm.Open(postgres.Open(dbConn.dsn), &gorm.Config{
			SkipDefaultTransaction: config.Manual,
		})
		return db, err
	}
	db, err := gorm.Open(postgres.Open(dbConn.dsn), &gorm.Config{})
	return db, err
}

func GetConnection() (*gorm.DB, error){
	session := NewPostgresSession(URI)
	config := &ConnectionConfig{Manual: false}
	dbConn, err := session.Connect(config)
	return dbConn, err
}
