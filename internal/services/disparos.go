package services

import (
	"fmt"

	"github.com/pedroafonso18/Disparador-GO/internal/database"
)

var limiteTemp int = 0

func disparos_evo(limite int) {
	campaigns, err := database.FetchCampanhas()
	if err != nil {
		fmt.Printf("error fetching campaigns: %v\n", err)
		return
	}

	for _, campaign := range campaigns {
		if limiteTemp == limite {
			return
		}

		fmt.Printf("Processing campaign: %s for number: %s\n", campaign.Campaign, campaign.Number)

		limiteTemp++
	}
}
