package main

import (
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

	sdk := synaps.InitSDK(apiKey)

	sessionID, err := sdk.Individual.Init(appID)
	if err != nil {
		log.Fatalf("failed to init session for app[%s]: %s", appID, err.Error())
	}

	_, err = sdk.Individual.Details(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s] and app[%s]: %s", sessionID, appID, err.Error())
	}

	// fmt.Println("session status: %s", details.Status)
}
