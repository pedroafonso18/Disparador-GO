package database

import (
	"context"
	"fmt"
	"log"
	"os"
)

func FetchConnections() {
	conn, err := GetConnection()
	if err != nil {
		log.Fatalf("Failed to find database")
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT name, instance_id, limite FROM instances WHERE active = true")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var instance_id string
		var limit uint
		if err := rows.Scan(&name, &instance_id, &limit); err != nil {
			fmt.Fprintf(os.Stderr, "Row scan with issues: %v\n ", err)
			continue
		}
		fmt.Println(name, instance_id, limit)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Rows iteration error: %v\n", err)
	}

}

func FetchCampanhas() {
	conn, err := GetConnection()
	if err != nil {
		log.Fatalf("Error at getting the campaigns.")
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT numero,campanha FROM campanhas WHERE disparado = FALSE")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error at getting query from db: %s\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var numero string
		var campanha string
		if err := rows.Scan(&numero, &campanha); err != nil {
			fmt.Fprintf(os.Stderr, "Error at reading line from campanhas: %s\n", err)
			continue
		}
		fmt.Printf(numero, campanha)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error at row iteration: %v\n", err)
	}
}
