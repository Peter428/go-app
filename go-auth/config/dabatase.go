package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql")

func DBConfig() (db *sql.DB, err error) {

	DBCONNECTION :="mysql"
	DBDATABASE :="go_auth"
	DBUSERNAME :="root"
	DBPASSWORD:="Asdf1234"

	db, err = sql.Open(DBCONNECTION, DBUSERNAME+":"+DBPASSWORD+"@/"+DBDATABASE)
	return
}
