# Synaps Go SDK

# Individual

The Synaps Individual Go SDK provides a convenient way to interact with the Synaps API, specifically tailored for individual sessions.  
Individual sessions, represent a Know Your Customer (KYC) session for a given user.  
This SDK enables you to initiate sessions, retrieve session details, and obtain information about different steps within a session (Liveness, Identity, Proof of address, etc.).

> For more details, please refer to the Synaps API documentation at [https://docs.synaps.io](https://docs.synaps.io).

## Installation

To use the Synaps Individual Go SDK, you can add it as a dependency in your project using `go get`:

```bash
go get github.com/synaps-hub/synaps-sdk-go/pkg/individual
```

## Requirements

Before you start using this SDK, ensure that you have the following:

- **Go Programming Language**: version 1.18 or higher.

- **Synaps API Key**: Your Synaps API key. You can find it on the [manager](https://manager-kyc.synaps.io) within the developer section of your app.

## Usage

The SDK facilitates the initiation of sessions, tracking user KYC progress, retrieving verification results, and handling events through webhooks.
This section provides an overview of the fundamental steps to integrate the SDK into your project and start leveraging its functionalities.

> A complete example can be found in the [examples/individual/main.go](https://github.com/synaps-hub/synaps-sdk-go/blob/main/examples/individual/main.go) file within the repository.

#### Import

```go
import (
	"github.com/synaps.io/synaps-sdk-go/pkg/individual"
)
```

#### Configuring client

Set the `SYNAPS_API_KEY` env variable to your api key value and create a new Synaps client from environment: 

```go
synapsClient := synaps.NewClientFromEnv()
```
> This will also check for `.env` file

Or create it from variables:
```go
synapsClient := synaps.NewClient("$YOUR_API_KEY")
```

#### Initialize session

Initialize a new session:

```go
initSessionRes, err := client.InitSession(synaps.InitSessionParams{})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID
```

Initialize a new session with an `alias`:

```go
initSessionRes, err := client.InitSession(synaps.InitSessionParams{Alias: "john-doe"})
if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID
```

Initialize a new session with `metadata`:

```go
initSessionRes, err := client.InitSession(synaps.InitSessionParams{Metadata: map[string]string{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john@doe.io",
}})

if err != nil {
	log.Fatalf("failed to init session: %s", err)
}
sessionID := initSessionRes.SessionID
```

#### Get session details

(Refer to the [documentation](https://docs.synaps.io/session#get-session-details) for details about the session details response)

```go
details, err := client.GetSessionDetails(sessionID)
if err != nil {
	log.Fatalf("failed to get details for session[%s]: %s", sessionID, err)
}

fmt.Printf("session status: %s\n", details.Session.Status)
```


#### Get step details
(Refer to the [documentation](https://docs.synaps.io/steps#get-step-details) for details about the step details response)

Get liveness step details using the `FindSessionStep` helper method:
```go
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
```

Get ID step details without helper method:
```go
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
```

Iterating through steps:
```go
...
var response any
var err error

for _, step := range details.Session.Steps {
	switch step.Type {
	case synaps.LivenessStep:
		response, err = client.GetStepLivenessDetails(sessionID, step.ID)
        // Do your stuff...
	case synaps.IDDocumentStep:
		response, err = client.GetStepIDDetails(sessionID, step.ID)
	case synaps.EmailStep:
		response, err = client.GetStepEmailDetails(sessionID, step.ID)
	case synaps.PhoneStep:
		response, err = client.GetStepPhoneDetails(sessionID, step.ID)
	case synaps.ProofOfAddressStep:
		response, err = client.GetStepProofOfAddressDetails(sessionID, step.ID)
    case synaps.AMLStep:
		response, err = client.GetStepAMLDetails(sessionID, step.ID)
	}

	if err != nil {
		log.Fatalf("failed to get step details for step [%s] and session[%s]: %s", step.Type, sessionID, err)
		continue
	}
	
    // Do your stuff...
}
```

### Webhooks

In order to receive webhooks, you'll need to create an endpoint that can receive and handle the webhook events. Below is an example of how to set up the necessary components.

> You can find the complete example in the [examples/individual/webhook/main.go](https://github.com/synaps-hub/synaps-sdk-go/blob/main/examples/individual/webhook/main.go) file within the repository.

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
Create your handler function for processing incoming requests:
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
	if r.URL.Query().Get("secret") != os.Getenv("SYNAPS_WEBHOOK_SECRET") {
		log.Printf("Error wrong webhook secret")
		http.Error(w, "Error invalid secret", http.StatusUnauthorized)
		return
	}

	handleEvent(payload)

	w.WriteHeader(http.StatusOK)
}
```

#### Handle event

Create a function to handle the webhook event:

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

#### Serve endpoint

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

Once done, add your endpoint URL to Synaps [manager](https://manager-kyc.synaps.io) (see [documentation](https://docs.synaps.io/quickstart#6-configure-webhooks) for guidance).

Congratulations, you're now all set!

Be sure not to overlook theses steps to ensure security:
- Verifying that the secret in the query parameters is matching the one given to you on the manager. This step ensures that you are exclusively receiving events from Synaps, as shown in the [example](https://github.com/synaps-hub/synaps-sdk-go/blob/main/examples/individual/webhook/main.go#L37) below.
- Utilizing HTTPS to establish a secure communication channel. This practice ensures the confidentiality and integrity of the data being exchanged.

## API Reference

For more details about the API, please refer to the [Synaps API Reference](https://docs.synaps.io/session).

# Corporate (Coming soon)
...

# License

This SDK is released under the [MIT License](LICENSE). Feel free to review the terms of the license in the provided [LICENSE](LICENSE) file.
