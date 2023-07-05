package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func main() {
	sdk := individual.NewClientFromEnv()

	sessionID, err := sdk.InitSession()
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}

	// Getting session sessionDetails

	sessionDetails, err := sdk.SessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", sessionDetails.Session.Status)

	// Method iterating over session steps to find the one with this type

	step, err := sessionDetails.GetSessionStep(individual.Liveness)
	if err != nil {
		log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
	}

	// Getting step details

	stepDetails, err := sdk.StepDetails(sessionID, step.ID)
	if err != nil {
		log.Fatalf("failed to get step details step [%s] and session[%s]: %s", step.Type, sessionID, err)
	}
	fmt.Printf("step details: %+v\n", stepDetails)

	// To get step relative data you can either do this

	switch stepDetails.Type {
	case "LIVENESS":
		stepData := individual.GetStepSpecificData[StepLiveness](stepDetails)
		fmt.Printf("Liveness file url: %s\n", stepData.Data.Liveness.File.URL)
	case "ID_DOCUMENT":
		stepData2 := individual.GetStepSpecificData[StepIdentity](stepDetails)
		fmt.Printf("Document firstname: %s\n", stepData2.Data.Fields.Firstname)
	}

	// Or simply use go classic json handling

	switch stepDetails.Type {
	case "LIVENESS":
		fmt.Printf("Liveness file url: %+v", stepDetails.Verification["liveness"].(map[string]any)["file"].(map[string]any)["url"].(string))
	case "ID_DOCUMENT":
		fmt.Printf("Document firstname: %s\n", stepDetails.Document["fields"].(map[string]any)["firstname"].(string))
	}
}
