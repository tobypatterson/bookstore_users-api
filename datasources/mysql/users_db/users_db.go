package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "gp2"
	mysql_users_password = "gp2"
	mysql_users_database = "gp2"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@/%s?charset=utf8",
		mysql_users_username,
		mysql_users_password,
		mysql_users_database,
	)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")
}
