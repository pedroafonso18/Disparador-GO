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
	var name string
	var instance_id string
	var limit uint
	err = conn.QueryRow(context.Background(), "SELECT name, instance_id, limite FROM instances WHERE active = true").Scan(&name, &instance_id, &limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(name, instance_id, limit)

}
