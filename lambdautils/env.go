package lambdautils

import (
	"fmt"
	"os"
)

// EnvQueueURL is the environment variable that contains the SQS queue URL.
// TODO: Can we use Golang build constraints to change this value?
const EnvQueueURL = "QUEUE_URL"

// Mustenv ensures an environment variable is set and panics if it is not.
func Mustenv(names ...string) {
	for _, name := range names {
		if os.Getenv(name) == "" {
			panic(fmt.Errorf("missing required environment variable: %v", name))
		}
	}
}

// QueueURL returns the SQS Queue URL set in the environment variable.
func QueueURL() string {
	return os.Getenv(EnvQueueURL)
}
