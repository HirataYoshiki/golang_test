package client

import (
	i "golang_test/dbinterface"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	Host string
	Port int
	User string
	Pass string
}

func (mysql MySQLClient) GetAll (response i.clientResponse, err error) {
	database, 
}