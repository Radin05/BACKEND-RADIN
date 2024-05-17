package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (db *sqlx.DB, err error) {
	var (
		host     = os.Getenv("DATABASE_HOST")
		port     = os.Getenv("DATABASE_PORT")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbname   = os.Getenv("DATABASE_NAME")
		mysqlUrl string
	)

	mysqlUrl = fmt.Sprintf("%s:%s@(%s:%s)/%s",
		user, password, host, port, dbname)

	db, err = sqlx.Open("mysql", mysqlUrl)
	if err != nil {
		log.Fatalf("Can't connect to database with err: %s", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Database not responding with err: %s\n", err)
		return
	}

	return
}
