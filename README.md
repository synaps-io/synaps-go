# Synaps Go SDK

# Individual 

The Individual Synaps Go SDK provides a convenient way to interact with the Synaps API for individual sessions.  
Individual sessions, represent a Know Your Customer (KYC) session for a given user. This SDK enables you to initiate sessions, retrieve session details, and obtain information about different steps within a session (Liveness, Identity, Proof of address, ...).  

> For more details the Synaps API documentation can be found at [https://docs.synaps.io](https://docs.synaps.io).

## Installation

To use the Synaps Go SDK, you can add it as a dependency in your project using `go get`:

```bash
go get github.com/synaps-hub/synaps-sdk-go/pkg/individual
```

## Requirements

Before you start using the Synaps Go SDK, ensure that you have the following:

- **Go Programming Language**: requires Go 1.19 or higher.

- **Synaps API Key**: To use the SDK, you need to have your Synaps API key. You can find it on the [manager](https://manager-kyc.synaps.io) in the developer section of your app.

## Usage

The SDK allows initiating sessions, tracking user KYC progress, retrieving verification results, and event handling using webhooks.  
This section provides an overview of the basic steps to integrate the SDK into your project and begin utilizing its features.  

> You can check the full example in the [exemples/individual/main.go](https://github.com/synaps-hub/synaps-sdk-go/blob/main/examples/individual/main.go) file within the repository.

#### Import

```go
import (
	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
)
```

#### Configuring client

Set the `SYNAPS_API_KEY` env variable to your api key and create a new Synaps client from environment: 

```go
synapsClient := synaps.NewClientFromEnv()
```
> This will also check for `.env` file

Or create it from variables:
```go
synapsClient := synaps.NewClient("API_KEY")
```

#### Init session

Initialize a new session:

```go
initSessionRes, err := synapsClient.InitSession(nil)
if err != nil {
    log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID
```

Initialize a new session with `alias`:

```go
alias := "john-doe"
initSessionRes, err := synapsClient.InitSession(&alias)
```

#### Get session details
(see [documentation](https://docs.synaps.io/session#get-session-details) for details about get session details response)

```go
sessionDetails, err := synapsClient.GetSessionDetails(sessionID)
if err != nil {
	log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
}
fmt.Printf("session status: %s\n", sessionDetails.Session.Status)
```


#### Get step details 
(see [documentation](https://docs.synaps.io/steps#get-step-details) for details about get step details response)

Get liveness step details using the `FindSessionStep` helper method:
```go
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
```

Get ID document step details without helper method:
```go
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
```

Iterating over steps:
```go
for _, step := range sessionDetails.Session.Steps {
	switch step.Type {
	case synaps.Liveness:
		_, _ = synapsClient.GetStepLivenessDetails(sessionID, step.ID)
	case synaps.IDDocument:
		_, _ = synapsClient.GetStepIDDocumentDetails(sessionID, step.ID)
	case synaps.Email:
		_, _ = synapsClient.GetStepEmailDetails(sessionID, step.ID)
	case synaps.Phone:
		_, _ = synapsClient.GetStepPhoneDetails(sessionID, step.ID)
	case synaps.ProofOfAddress:
		_, _ = synapsClient.GetStepProofOfAddressDetails(sessionID, step.ID)
	}
}
```

### Webhooks

In order to receive webhooks, you'll need to create an endpoint that can receive and handle the webhook events. Below is an example of how to set up the necessary components.

> You can check the full example in the [exemples/individual/webhook/main.go](https://github.com/synaps-hub/synaps-sdk-go/blob/main/examples/individual/webhook/main.go) file within the repository.

#### Import
```go
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
)
```

#### Handle webhook
Create your handler function for processing incoming request:
```go
func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

        // Unmarshaling body
	var payload synaps.WebhookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error unmarshaling request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

        // Checking for secret
	if r.URL.Query().Get("secret") != os.Getenv("secret") {
		log.Printf("Error wrong webhook secret")
		http.Error(w, "Error invalid secret", http.StatusUnauthorized)
		return
	}

	handleEvent(payload)

	w.WriteHeader(http.StatusOK)
}
```

#### Handle event
Create the function that handles event itself:
```go
func handleEvent(payload synaps.WebhookPayload) {
	switch payload.Status {
	case synaps.EventApproved:
		log.Printf("Received event: APPROVED")
	case synaps.EventRejected:
		log.Printf("Received event: REJECTED")
        // ...
	}
}
```

Serve your endpoint:
```go
func main() {
	_, ok := os.LookupEnv("SYNAPS_WEBHOOK_SECRET")
	if !ok {
		log.Fatalf("Error missing webhook secret")
	}

	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("Webhook server listening on port 80...")
	http.ListenAndServe(":80", nil)
}
```

> Ensure that your endpoint is reachable from the internet so webhook server can reach it

Once its done you can add your endpoint URL to synaps [manager](https://manager-kyc.synaps.io) (see [documentation](https://docs.synaps.io/quickstart#6-configure-webhooks) for guidance)  

Congratulations, you're now all set!

Remember to keep it secure by:  
- Verifying that the secret in the query parameters is matching the one given to you on the manager. This step ensures that you are exclusively receiving events from Synaps, as shown in the [example](https://github.com/synaps-hub/synaps-sdk-go/tree/refactor/exemple-and-error-handling#handle-webhook) below.
- Utilizing HTTPS to establish a secure communication channel. This practice ensures the confidentiality and integrity of the data being exchanged.

## API Reference

For more details on the API, please refer to the [Synaps API Reference](https://docs.synaps.io/session).

# Corporate (Coming soon)
...

# License

This SDK is released under the [MIT License](LICENSE). Feel free to review the terms of the license in the provided [LICENSE](LICENSE) file.
