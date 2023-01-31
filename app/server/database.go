package server

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

const (
	host     = "localhost"
	user     = "root"
	password = "pass"
	dbname   = "sitegolang"
)

func InitDb() (err error) {

	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		user, password, host, dbname)

	Db, err = sqlx.Connect("mysql", psqlInfo)
	if err != nil {
		return
	}

	err = Db.Ping()
	return
}
