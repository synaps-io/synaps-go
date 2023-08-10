package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
	"github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func main() {
	synapsClient := individual.NewClientFromEnv()

	req := models.InitSessionRequest{Alias: "username", Metadata: map[string]string{"email": "john.doe@gmail.com"}}
	initSessionRes, err := synapsClient.InitSession(req)
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}
	sessionID := initSessionRes.SessionID

	// Getting session details

	sessionDetails, err := synapsClient.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}
	fmt.Printf("session status: %s\n", sessionDetails.Session.Status)

	// Getting liveness step details with FindSessionStep helper method

	{
		livenessStep, err := sessionDetails.FindSessionStep(models.LivenessStep)
		if err != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		livenessStepDetails, err := synapsClient.GetStepLivenessDetails(sessionID, livenessStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
		}

		fmt.Printf("Liveness step status: %s\n", livenessStep.Status)

		if livenessStep.Status == models.Approved {
			fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
		}

		if livenessStep.Status == models.Rejected {
			fmt.Printf("Liveness reject reason: %s\n", livenessStepDetails.Reason.Message)
		}
	}

	// Getting id document step details without helper method

	{
		var IDDocumentStep *models.Step
		for _, step := range sessionDetails.Session.Steps {
			if step.Type == models.IDDocumentStep {
				IDDocumentStep = &step
				break
			}
		}

		if IDDocumentStep == nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		idDocumentStepDetails, err := synapsClient.GetStepIDDocumentDetails(sessionID, IDDocumentStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", IDDocumentStep.Type, sessionID, err)
		}

		fmt.Printf("ID Document step status: %s\n", idDocumentStepDetails.Status)

		if idDocumentStepDetails.Status == models.Pending || idDocumentStepDetails.Status == models.Approved {
			fmt.Printf("ID Document firstname: %s\n", idDocumentStepDetails.Document.Fields.Firstname)
		}

	}

	// Iterating over steps

	{
		for _, step := range sessionDetails.Session.Steps {
			switch step.Type {
			case models.LivenessStep:
				_, _ = synapsClient.GetStepLivenessDetails(sessionID, step.ID)
			case models.IDDocumentStep:
				_, _ = synapsClient.GetStepIDDocumentDetails(sessionID, step.ID)
			case models.EmailStep:
				_, _ = synapsClient.GetStepEmailDetails(sessionID, step.ID)
			case models.PhoneStep:
				_, _ = synapsClient.GetStepPhoneDetails(sessionID, step.ID)
			case models.ProofOfAddressStep:
				_, _ = synapsClient.GetStepProofOfAddressDetails(sessionID, step.ID)
			}
		}
	}
}
