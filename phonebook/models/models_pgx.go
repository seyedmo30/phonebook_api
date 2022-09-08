package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn
var err error

func Setup_v2() {
	conn, err = pgx.Connect(context.Background(), "postgres://postgres:salam@localhost:5432/phonebook")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}

}
