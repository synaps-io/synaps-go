package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
)

func main() {
	synapsClient := synaps.NewClientFromEnv()

	alias := "john-doe"
	initSessionRes, err := synapsClient.InitSession(&alias)
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
		livenessStep, err := sessionDetails.FindSessionStep(synaps.LivenessStep)
		if err != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		livenessStepDetails, err := synapsClient.GetStepLivenessDetails(sessionID, livenessStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
		}

		fmt.Printf("Liveness step status: %s\n", livenessStep.Status)

		if livenessStep.Status == synaps.StatusApproved {
			fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
		}

		if livenessStep.Status == synaps.StatusRejected {
			fmt.Printf("Liveness reject reason: %s\n", livenessStepDetails.Reason.Message)
		}
	}

	// Getting id document step details without helper method

	{
		var IDDocumentStep *synaps.Step
		for _, step := range sessionDetails.Session.Steps {
			if step.Type == synaps.IDDocumentStep {
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

		if idDocumentStepDetails.Status == synaps.StatusPending || idDocumentStepDetails.Status == synaps.StatusApproved {
			fmt.Printf("ID Document firstname: %s\n", idDocumentStepDetails.Document.Fields.Firstname)
		}

	}

	// Iterating over steps

	{
		for _, step := range sessionDetails.Session.Steps {
			switch step.Type {
			case synaps.LivenessStep:
				_, _ = synapsClient.GetStepLivenessDetails(sessionID, step.ID)
			case synaps.IDDocumentStep:
				_, _ = synapsClient.GetStepIDDocumentDetails(sessionID, step.ID)
			case synaps.EmailStep:
				_, _ = synapsClient.GetStepEmailDetails(sessionID, step.ID)
			case synaps.PhoneStep:
				_, _ = synapsClient.GetStepPhoneDetails(sessionID, step.ID)
			case synaps.ProofOfAddressStep:
				_, _ = synapsClient.GetStepProofOfAddressDetails(sessionID, step.ID)
			}
		}
	}
}
