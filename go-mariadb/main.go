package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func QueryDB(db *sql.DB, cmd string, args ...any) {
	if args == nil {
		log.Print("!!!!!!!")
	}
	rows, err := db.Query(cmd, args...)
	if err != nil {
		log.Print("Query error:")
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("id %d name is %s", id, name)
	}
}

func main() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "root:stex@/xxx")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Printf("Version: %s", version)

	// Query
	QueryDB(db, "SELECT * FROM controller")

	log.Print("=== Second ===")
	QueryDB(db, `Select * from controller where name=?`, "L85E")

	// Transaction
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "INSERT INTO controller(id, name) VALUES(?, ?),(?,?)", 123, "A", 456, "B")
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("Insert failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("Insert failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Insert success")
	}
}
