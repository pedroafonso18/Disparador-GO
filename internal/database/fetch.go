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
