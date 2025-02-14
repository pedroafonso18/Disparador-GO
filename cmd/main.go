package main

import (
	"fmt"

	"github.com/pedroafonso18/Disparador-GO/internal/config"
	"github.com/pedroafonso18/Disparador-GO/internal/database"
	"github.com/pedroafonso18/Disparador-GO/internal/services"
)

func main() {
	if services.ReturnTime() {
		config.Load()
		database.FetchConnections()
		database.FetchCampanhas()
		services.ReturnTime()
	} else {
		fmt.Println("Espere um instante e tente novamente.")
	}
}
