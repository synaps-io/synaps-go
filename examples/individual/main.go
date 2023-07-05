package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

func main() {
	synapsClient := individual.NewClientFromEnv()

	sessionID, err := synapsClient.InitSession()
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}

	// Getting session details

	sessionDetails, err := synapsClient.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", sessionDetails.Session.Status)

	// Getting liveness step details with GetSessionStep helper method

	func() {
		livenessStep, err := sessionDetails.FindSessionStep(individual.Liveness)
		if err != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		stepDetails, err := synapsClient.GetStepDetails(sessionID, livenessStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
		}

		livenessStepDetails := stepDetails.(LivenessStepDetails)

		fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
	}()

	// Getting id document step details without using helper method

	func() {
		var IDDocumentStep *Step
		for _, step := range sessionDetails.Session.Steps {
			if step.Type == individual.IDDocument {
				IDDocumentStep = &step
				break
			}
		}

		if IDDocumentStep != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}

		stepDetails, err := synapsClient.GetStepDetails(sessionID, IDDocumentStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details step [%s] and session[%s]: %s", IDDocumentStep.Type, sessionID, err)
		}

		idDocumentStepDetails := stepDetails.(IDDocumentStepDetails)

		fmt.Printf("ID Document firstname: %s\n", idDocumentStepDetails.Document.Fields.Firstname)
	}()

	// Iterating over steps

	func() {
		for _, step := range sessionDetails.Session.Steps {
			stepDetails, err := synapsClient.GetStepDetails(sessionID, step.ID)
			if err != nil {
				return
			}

			switch step.Type {
			case individual.Liveness:
				_ = stepDetails.(LivenessStepDetails)
			case individual.IDDocument:
				_ = stepDetails.(IDDocumentStepDetails)
			case individual.Email:
				_ = stepDetails.(EmailStepDetails)
			case individual.Phone:
				_ = stepDetails.(PhoneStepDetails)
			case individual.ProofOfAddress:
				_ = stepDetails.(ProofOfAddressStepDetails)
			}
		}
	}()
}
