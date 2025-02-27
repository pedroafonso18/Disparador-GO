package main

import (
	"fmt"
	"time"

	"math/rand"

	"github.com/pedroafonso18/Disparador-GO/internal/config"
	"github.com/pedroafonso18/Disparador-GO/internal/services"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		config.Load()
		if services.ReturnTime() {
			services.Disparos()
		} else {
			fmt.Println("Fora do horário de funcionamento, espere e tente novamente.")
		}
	}
}
