package main

import (
	"fmt"
	"log"
	"os"

	"github.com/synaps.io/synaps-sdk-go"
)

func main() {
	apiKey := os.Getenv("SYNAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("SYNAPS_API_KEY is not set")
	}

	appID := os.Getenv("SYNAPS_APP_ID")
	if appID == "" {
		log.Fatal("SYNAPS_APP_ID is not set")
	}

	sdk := synaps.NewIndividualSDK(apiKey)

	sessionID, err := sdk.InitIndividual()
	if err != nil {
		log.Fatalf("failed to init session for app[%s]: %s", appID, err.Error())
	}

	details, err := sdk.Details(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s] and app[%s]: %s", sessionID, appID, err.Error())
	}

	fmt.Printf("session status: %s\n", details.Status)

	overview, err := sdk.Details(sessionID)
	if err != nil {
		log.Fatalf("failed to get overview for session[%s] and app[%s]: %s", sessionID, appID, err.Error())
	}

	fmt.Printf("overview: %+v\n", overview)
}
