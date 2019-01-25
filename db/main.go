package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	con := "HOSTNAME=127.0.0.1;DATABASE=SDMWS;PORT=50000;UID=db2inst1;PWD=superGeheimesPasswort"
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	result, err := db.ExecContext(ctx,
		"INSERT INTO CUSTOMER (id, firstname) VALUES (?, ?)",
		1,
		"gopher")

	if err != nil {
		log.Fatal(err)
	}
	lastInsertID, err := result.LastInsertId()
	rowsAffected, err := result.RowsAffected()

	fmt.Println("Insert", lastInsertID, rowsAffected)

	rows, err := db.QueryContext(ctx, "SELECT firstname FROM customer where id = ?", 2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var firstname string
		if err := rows.Scan(&firstname); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s", firstname)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	db.Close()

}
