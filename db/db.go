package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/koesterjannik/starter/logger"
)

var Db *pgx.Conn

func ConnectToDb() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == "" {
		logger.Logger.Error("Environment variable DATABASE_URL not set.")
		os.Exit(1)
	}
	logger.Logger.Info("Connecting to database with url: " + database_url)
	Db, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer Db.Close(context.Background())

	logger.Logger.Info("Connected to database")
	/*var name string
	var weight int64
	err = Db.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(name, weight)*/
}
