package services

import (
	"fmt"
	"sync"
	"time"

	"math/rand"

	"github.com/pedroafonso18/Disparador-GO/internal/api"
	"github.com/pedroafonso18/Disparador-GO/internal/database"
)

func Disparos() {
	fmt.Println("Starting Disparos operation")

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

	var wg sync.WaitGroup

	for _, connection := range connections {
		wg.Add(1)
		go func(conn database.Instance) {
			defer wg.Done()
			messagesSent := uint(0)

			for _, campaign := range campaigns {
				if messagesSent >= conn.Limit {
					fmt.Printf("Reached message limit for connection: %s\n", conn.Name)
					return
				}

				delay := rand.Intn(5) + 3

				fmt.Printf("Sending message for campaign: %s to number: %s using %s\n",
					campaign.Campaign, campaign.Number, conn.Name)

				var err error
				if conn.IsEvo {
					err = api.SendMessageEvo(campaign.Number, conn.Name, templates)
				} else {
					err = api.SendMessageWuz(campaign.Number, templates, conn.InstanceID)
				}

				if err != nil {
					fmt.Printf("Error sending message through %s: %s\n", conn.Name, err)
					continue
				}

				database.UpdateDisparados(campaign.Number)
				database.InsertLog(campaign.Number, conn.Name, templates, campaign.Campaign)
				messagesSent++

				fmt.Printf("Connection %s waiting %d seconds before next message\n", conn.Name, delay)
				time.Sleep(time.Duration(delay) * time.Second)
			}
		}(connection)
	}

	wg.Wait()
	fmt.Println("Disparos operation completed")
}
