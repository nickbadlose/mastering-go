package main

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
	"strconv"

	_ "github.com/jackc/pgx"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 5 {
		panic("Please provide: hostname port username password db")
	}

	host := arguments[0]
	p := arguments[1]
	user := arguments[2]
	pass := arguments[3]
	database := arguments[4]

	port, err := strconv.Atoi(p)
	if err != nil {
		panic(fmt.Sprintf("Not a valid port number: %s", err))
	}

	db, err := pgx.Connect(pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		Database: database,
		User:     user,
		Password: pass,
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "datname" FROM "pg_database" WHERE datistemplate = false`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		rErr := rows.Scan(&name)
		if rErr != nil {
			panic(rErr)
		}

		fmt.Println("*", name)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}

	// Get all tables from __current__ database
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
	rows, err = db.Query(query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
