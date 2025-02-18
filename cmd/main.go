package main

import (
	"fmt"

	"github.com/pedroafonso18/Disparador-GO/internal/config"
	"github.com/pedroafonso18/Disparador-GO/internal/services"
)

func main() {
	for 1 > 0 {
		config.Load()
		if services.ReturnTime() {
			services.Disparos()
		} else {
			fmt.Println("Espere um instante e tente novamente.")
		}
	}
}
