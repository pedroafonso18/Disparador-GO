package main

import (
	"github.com/pedroafonso18/Disparador-GO/internal/config"
	"github.com/pedroafonso18/Disparador-GO/internal/database"
)

func main() {
	config.Load()
	database.FetchConnections()
}
