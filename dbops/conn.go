package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:abc123456@tcp(10.211.55.5:3306)/test?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}