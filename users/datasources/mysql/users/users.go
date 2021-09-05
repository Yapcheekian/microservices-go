package users

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "MYSQL_USER"
	mysqlPassword = "MYSQL_PASS"
	mysqlHost     = "MYSQL_HOST"
	mysqlDatabase = "MYSQL_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	host     = os.Getenv(mysqlHost)
	database = os.Getenv(mysqlDatabase)
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, database)
	var err error
	Client, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
