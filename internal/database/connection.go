package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/pedroafonso18/Disparador-GO/internal/config"
)

var (
	conn     *pgx.Conn
	connErr  error
	connOnce sync.Once
)

func GetConnection() (*pgx.Conn, error) {
	connOnce.Do(func() {
		conn, connErr = pgx.Connect(context.Background(), config.DBURL)
		if connErr != nil {
			fmt.Printf("Unable to connect do DB")
		}
	})
	return conn, connErr

}
