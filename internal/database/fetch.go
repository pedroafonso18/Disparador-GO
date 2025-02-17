package database

import (
	"context"
	"fmt"
	"os"
)

type Campaign struct {
	Number   string
	Campaign string
}

type Instance struct {
	Name       string
	InstanceID string
	Limit      uint
	IsEvo      bool
}

func FetchConnections() ([]Instance, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to find database: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT name, instance_id, limite, is_evo FROM instances WHERE active = true")
	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	var instances []Instance
	for rows.Next() {
		var inst Instance
		if err := rows.Scan(&inst.Name, &inst.InstanceID, &inst.Limit, &inst.IsEvo); err != nil {
			fmt.Fprintf(os.Stderr, "Row scan with issues: %v\n", err)
			continue
		}
		instances = append(instances, inst)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return instances, nil
}

func FetchCampanhas() ([]Campaign, error) {
	conn, err := GetConnection()
	if err != nil {
		return nil, fmt.Errorf("error getting campaigns: %v", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT numero, campanha FROM campanhas WHERE disparado = FALSE")
	if err != nil {
		return nil, fmt.Errorf("error getting query from db: %v", err)
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var camp Campaign
		if err := rows.Scan(&camp.Number, &camp.Campaign); err != nil {
			fmt.Fprintf(os.Stderr, "Error at reading line from campanhas: %v\n", err)
			continue
		}
		campaigns = append(campaigns, camp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error at row iteration: %v", err)
	}

	return campaigns, nil
}
