package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func CreateTables() error {
	stmtCr, err := dbConn.Prepare(`CREATE TABLE users (
		id int unique not null primary ,
		login_name not null varchar(64),
		pwd NOT NULL TEXT,
)`)
}