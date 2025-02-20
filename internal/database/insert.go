package database

import (
	"context"
	"fmt"
)

func UpdateDisparados(num string) {
	conn, err := GetConnection()
	if err != nil {
		fmt.Println("Error getting connection:", err)
		return
	}
	defer conn.Close(context.Background())

	query := "UPDATE campanhas SET disparado = true WHERE numero = $1"
	result, err := conn.Exec(context.Background(), query, num)
	if err != nil {
		fmt.Println("Error executing update:", err)
		return
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Update successful, %d rows affected.\n", rowsAffected)
}

func InsertLog(num string, conexao string, template string, nome_campanha string) {
	conn, err := GetConnection()
	if err != nil {
		fmt.Println("Error connecting to insert log:", err)
		return
	}
	defer conn.Close(context.Background())

	query := "INSERT INTO logs (num, conexao, template, nome_campanha) VALUES ($1, $2, $3, $4)"
	result, err := conn.Exec(context.Background(), query, num, conexao, template, nome_campanha)
	if err != nil {
		fmt.Println("Error executing insert:", err)
		return
	}

	query2 := "UPDATE instances SET sent = sent + 1 WHERE name = $1"
	_, err = conn.Exec(context.Background(), query2, conexao)
	if err != nil {
		fmt.Println("Error executing update on instances:", err)
		return
	}

	rowsAffected := result.RowsAffected()

	fmt.Printf("Insert successful, %d rows affected.\n", rowsAffected)
}
