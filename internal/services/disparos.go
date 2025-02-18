package services

import (
	"fmt"

	"github.com/pedroafonso18/Disparador-GO/internal/api"
	"github.com/pedroafonso18/Disparador-GO/internal/database"
)

func Disparos() {
	templates, err := database.FetchTemplateText()
	if err != nil {
		fmt.Printf("error getting templates: %s\n", err)
		return
	}

	connections, err := database.FetchConnections()
	if err != nil {
		fmt.Printf("error getting connections: %s\n", err)
		return
	}

	campaigns, err := database.FetchCampanhas()
	if err != nil {
		fmt.Printf("error getting campaigns: %s\n", err)
		return
	}

	for _, connection := range connections {
		messagesSent := uint(0)

		for _, campaign := range campaigns {
			if messagesSent >= connection.Limit {
				break
			}

			if connection.IsEvo {
				err = api.SendMessageEvo(campaign.Number, connection.Name, templates)
			} else {
				err = api.SendMessageWuz(campaign.Number, templates, connection.InstanceID)
			}

			if err != nil {
				fmt.Printf("error sending message: %s\n", err)
				continue
			}

			messagesSent++
		}
	}
}
