package main

import (
	"fmt"

	"github.com/pedroafonso18/Disparador-GO/internal/config"
	"github.com/pedroafonso18/Disparador-GO/internal/services"
)

func main() {
	for {
		config.Load()
		if services.ReturnTime() {
			services.Disparos()
		} else {
			fmt.Println("Fora do horário de funcionamento, espere e tente novamente.")
		}
	}
}
