package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/synaps.io/synaps-sdk-go/pkg/corporate"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("SYNAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("SYNAPS_API_KEY is not set")
	}

	appID := os.Getenv("SYNAPS_APP_ID")
	if appID == "" {
		log.Fatal("SYNAPS_APP_ID is not set")
	}

	sdk := corporate.NewClient(apiKey)

	sessionID, err := sdk.Init()
	if err != nil {
		log.Fatalf("failed to init session for app[%s]: %s", appID, err.Error())
	}

	details, err := sdk.Details(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s] and app[%s]: %s", sessionID, appID, err.Error())
	}

	fmt.Printf("session status: %s\n", details.Status)

	overview, err := sdk.Overview(sessionID)
	if err != nil {
		log.Fatalf("failed to get overview for session[%s] and app[%s]: %s", sessionID, appID, err.Error())
	}

	fmt.Printf("overview: %+v\n", overview)
}
