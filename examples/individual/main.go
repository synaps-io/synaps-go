package main

import (
	"fmt"
	"log"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
)

func main() {
	client := synaps.NewClientFromEnv()

	alias := "john-doe"
	initSessionRes, err := client.InitSession(alias)
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}
	sessionID := initSessionRes.SessionID

	// Getting session details
	details, err := client.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}

	fmt.Printf("session status: %s\n", details.Session.Status)

	processLiveness(client, details)

	processID(client, details)

	processSteps(client, details)
}

// Getting liveness step details with FindSessionStep helper method
func processLiveness(client *synaps.Client, details synaps.SessionDetailsResponse) {
	sessionID := details.Session.ID

	livenessStep, err := details.FindSessionStep(synaps.LivenessStep)
	if err != nil {
		log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
	}

	livenessStepDetails, err := client.GetStepLivenessDetails(sessionID, livenessStep.ID)
	if err != nil {
		log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
	}

	fmt.Printf("Liveness step status: %s\n", livenessStep.Status)

	switch livenessStep.Status {
	case synaps.StatusApproved:
		fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
	case synaps.StatusRejected:
		fmt.Printf("Liveness reject reason: %s\n", livenessStepDetails.Reason.Message)
	default:
		fmt.Printf("Liveness step is not finished yet\n")
	}
}

// Getting id document step details without helper method
func processID(client *synaps.Client, details synaps.SessionDetailsResponse) {
	var IDStep *synaps.Step
	for _, step := range details.Session.Steps {
		if step.Type == synaps.IDDocumentStep {
			IDStep = &step
			break
		}
	}

	sessionID := details.Session.ID
	if IDStep == nil {
		log.Fatalf("failed to get step for session[%s]", sessionID)
	}

	IDStepDetails, err := client.GetStepIDDetails(sessionID, IDStep.ID)
	if err != nil {
		log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", IDStep.Type, sessionID, err)
	}

	fmt.Printf("ID step status: %s\n", IDStepDetails.Status)

	if IDStepDetails.Status == synaps.StatusPending || IDStepDetails.Status == synaps.StatusApproved {
		fmt.Printf("ID Document firstname: %s\n", IDStepDetails.Document.Fields["FIRSTNAME"])
	}
}

// Iterating over steps
func processSteps(client *synaps.Client, details synaps.SessionDetailsResponse) {
	sessionID := details.Session.ID

	var response any
	var err error
	for _, step := range details.Session.Steps {
		switch step.Type {
		case synaps.LivenessStep:
			response, err = client.GetStepLivenessDetails(sessionID, step.ID)
		case synaps.IDDocumentStep:
			response, err = client.GetStepIDDetails(sessionID, step.ID)
		case synaps.EmailStep:
			response, err = client.GetStepEmailDetails(sessionID, step.ID)
		case synaps.PhoneStep:
			response, err = client.GetStepPhoneDetails(sessionID, step.ID)
		case synaps.ProofOfAddressStep:
			response, err = client.GetStepProofOfAddressDetails(sessionID, step.ID)
		}

		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
			continue
		}
		fmt.Printf("Response is:\n%+v\n", response)
	}
}
