package services

import (
	"fmt"
	"time"

	"math/rand/v2"

	"github.com/pedroafonso18/Disparador-GO/internal/api"
	"github.com/pedroafonso18/Disparador-GO/internal/database"
)

func Disparos() {
	fmt.Println("Starting Disparos operation")
	rand := rand.IntN(8-3) + 3

	fmt.Println("Fetching template texts")
	templates, err := database.FetchTemplateText()
	if err != nil {
		fmt.Printf("Error getting templates: %s\n", err)
		return
	}
	fmt.Println("Successfully fetched template texts")

	fmt.Println("Fetching active campaign configuration")
	config, err := database.FetchActiveCampanha()
	if err != nil {
		fmt.Printf("Error getting active campaign configuration: %s\n", err)
		return
	}
	fmt.Println("Successfully fetched active campaign configuration")

	fmt.Println("Fetching connections")
	connections, err := database.FetchConnections()
	if err != nil {
		fmt.Printf("Error getting connections: %s\n", err)
		return
	}
	fmt.Println("Successfully fetched connections")

	fmt.Println("Fetching campaigns")
	campaigns, err := database.FetchCampanhas(config)
	if err != nil {
		fmt.Printf("Error getting campaigns: %s\n", err)
		return
	}
	fmt.Println("Successfully fetched campaigns")

	for _, connection := range connections {
		fmt.Printf("Processing connection: %s\n", connection.Name)
		messagesSent := uint(0)

		for _, campaign := range campaigns {
			if messagesSent >= connection.Limit {
				fmt.Printf("Reached message limit for connection: %s\n", connection.Name)
				break
			}

			fmt.Printf("Sending message for campaign: %s to number: %s\n", campaign.Campaign, campaign.Number)
			if connection.IsEvo {
				err = api.SendMessageEvo(campaign.Number, connection.Name, templates)
				database.UpdateDisparados(campaign.Number)
				database.InsertLog(campaign.Number, connection.Name, templates, campaign.Campaign)
				time.Sleep(time.Duration(rand) * time.Second)
			} else {
				err = api.SendMessageWuz(campaign.Number, templates, connection.InstanceID)
				database.UpdateDisparados(campaign.Number)
				database.InsertLog(campaign.Number, connection.Name, templates, campaign.Campaign)
				time.Sleep(time.Duration(rand) * time.Second)
			}

			if err != nil {
				fmt.Printf("Error sending message: %s\n", err)
				continue
			}

			fmt.Printf("Successfully sent message for campaign: %s to number: %s\n", campaign.Campaign, campaign.Number)
			messagesSent++
		}
	}

	fmt.Println("Disparos operation completed")
}
