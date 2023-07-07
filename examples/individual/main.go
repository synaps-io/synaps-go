package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func main() {
	synapsClient := individual.NewClientFromEnv()

	req := InitSessionRequest{Alias: "12345"}

	initSessionRes, err := synapsClient.InitSession(req)
	sessionID := initSessionRes.SessionID

	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}

	// Getting session details

	sessionDetails, err := synapsClient.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", sessionDetails.Session.Status)

	// Getting liveness step details with FindSessionStep helper method

	func() {
		livenessStep, err := sessionDetails.FindSessionStep(individual.Liveness)
		if err != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		livenessStepDetails, err := synapsClient.GetStepLivenessDetails(sessionID, livenessStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
		}

		fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
	}()

	// Getting id document step details without helper method

	func() {
		var IDDocumentStep *Step
		for _, step := range sessionDetails.Session.Steps {
			if step.Type == individual.IDDocument {
				IDDocumentStep = &step
				break
			}
		}

		if IDDocumentStep == nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		idDocumentStepDetails, err := synapsClient.GetStepIDDocumentDetails(sessionID, IDDocumentStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details step [%s] and session[%s]: %s", IDDocumentStep.Type, sessionID, err)
		}

		fmt.Printf("ID Document firstname: %s\n", idDocumentStepDetails.Document.Fields.Firstname)
	}()

	// Iterating over steps

	func() {
		for _, step := range sessionDetails.Session.Steps {
			switch step.Type {
			case individual.Liveness:
				_, _ = synapsClient.GetStepLivenessDetails(sessionID, step.ID)
			case individual.IDDocument:
				_, _ = synapsClient.GetStepIDDocumentDetails(sessionID, step.ID)
			case individual.Email:
				_, _ = synapsClient.GetStepEmailDetails(sessionID, step.ID)
			case individual.Phone:
				_, _ = synapsClient.GetStepPhoneDetails(sessionID, step.ID)
			case individual.ProofOfAddress:
				_, _ = synapsClient.GetStepProofOfAddressDetails(sessionID, step.ID)
			}
		}
	}()
}
