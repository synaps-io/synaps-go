# Synaps Go SDK

The Individual Synaps Go SDK provides a convenient way to interact with the Synaps API for individual sessions. Individual sessions, in the context of this SDK, represent a Know Your Customer (KYC) session for a given user. This SDK enables you to initiate sessions, retrieve session details, and obtain information about different steps within a session (Liveness, Identity, ...).

For more details the Synaps API documentation can be found at [https://docs.synaps.io](https://docs.synaps.io).

## Installation

To use the Synaps Go SDK, you can add it as a dependency in your project using `go get`:

```bash
go get github.com/synaps.io/synaps-sdk-go/pkg/individual
```

## Requirements

Before you start using the Synaps Go SDK, ensure that you have the following:

- **Go Programming Language**: The Synaps Go SDK requires Go 1.19 or higher.

- **Synaps API Credentials**: To use the SDK, you need to have your Synaps API credentials, including the API key or other authentication details. Theses can be found on the manager app [https://manager-kyc.synaps.io](https://manager-kyc.synaps.io)

## Usage

Import the required packages and use the SDK to interact with the Synaps API.

### Imports 

Import the sdk and models

```go

import (
	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
	. "github.com/synaps.io/synaps-sdk-go/pkg/individual/models"
)

```

### Configuring client

Create a new Synaps client from environment variables

```go
	synapsClient := individual.NewClientFromEnv()

```
Or create it from variables

```go
	synapsClient := individual.NewClient("BASE_URL", "API_KEY")
```

### Init session

Initialize a new session with `alias` 

```go
	req := InitSessionRequest{Alias: "12345"}
	initSessionRes, err := synapsClient.InitSession(req)
	if err != nil {
		log.Fatalf("failed to init session: %s", err)
	}
	sessionID := initSessionRes.SessionID

```

### Get session details

Get session details (see [https://docs.synaps.io/session#get-session-details] for details about get session details response)
```go
	sessionDetails, err := synapsClient.GetSessionDetails(sessionID)
	if err != nil {
		log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
	}
	fmt.Printf("session status: %s\n", sessionDetails.Session.Status)

```


### Get step details 

Get liveness step details using the `FindSessionStep` helper method (see [https://docs.synaps.io/steps#get-step-details] for details about get step details response)
```go
	{
		livenessStep, err := sessionDetails.FindSessionStep(individual.Liveness)
		if err != nil {
			log.Fatalf("failed to get step for session[%s]: %s", sessionID, err)
		}
		livenessStepDetails, err := synapsClient.GetStepLivenessDetails(sessionID, livenessStep.ID)
		if err != nil {
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", livenessStep.Type, sessionID, err)
		}
		fmt.Printf("Liveness step status: %s\n", livenessStep.Status)
		if livenessStep.Status == string(individual.Approved) {
			fmt.Printf("Liveness file url: %s\n", livenessStepDetails.Verification.Liveness.File.URL)
		}
	}
```

Get ID document step details without helper method
```go
	{
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
			log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", IDDocumentStep.Type, sessionID, err)
		}
		fmt.Printf("ID Document step status: %s\n", idDocumentStepDetails.Status)
		if idDocumentStepDetails.Status == string(individual.Pending) || idDocumentStepDetails.Status == string(individual.Approved) {
			fmt.Printf("ID Document firstname: %s\n", idDocumentStepDetails.Document.Fields.Firstname)
		}
	}
```

Iterating over steps 
```go
	{
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
	}
}

```

> You can check the full example in the [exemples/individual/main.go](https://github.com/synaps-hub/synaps/blob/main/exemples/individual/main.go) file within the repository.

## API Reference

For more details on the API, please refer to the [Synaps API Reference](https://docs.synaps.io/session).

