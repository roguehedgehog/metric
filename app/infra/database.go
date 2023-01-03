package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var PrimaryDb *sql.DB

func Healthy() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := PrimaryDb.PingContext(ctx); err != nil {
		log.Print(err)
		return false
	}

	return true
}

func init() {
	_, disabled := os.LookupEnv("METRIC_DB_DISABLED")
	if disabled {
		return
	}
	var err error
	PrimaryDb, err = getDatabaseConn(
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	if err != nil || !Healthy() {
		log.Fatal(err)
	}
}

func getDatabaseConn(host, port, username, password, database string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	log.Print(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
