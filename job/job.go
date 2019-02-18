package job

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var jobs map[string]NewJob

// LogJobName is LogJob's unique identifier.
const LogJobName = "log"

// Job is the interfce implemented by jobs.
type Job interface {
	Do(context.Context) (context.Context, error)
}

// NewJob is a function definition for functions that return initialized Jobs.
type NewJob = func() Job

// LogJob is a job that simply prints a log message.
type LogJob struct{}

// NewLogJob returns a LogJob.
func NewLogJob() Job {
	return LogJob{}
}

// Do implements Job.Do by printing a log message.
func (j LogJob) Do(ctx context.Context) (context.Context, error) {
	log.Println("job complete")
	return ctx, nil
}

// Message models a job message sent to the queue.
type Message struct {
	Job string `json:"job"`
}

// NewMessage returns an initialized Message.
func NewMessage(job string) Message {
	return Message{Job: job}
}

// String implementes Stringer by marshaling the message to JSON.
func (m Message) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Errorf("error encoding message: %v", err))
	}
	return string(b)
}

// Do parses the SQS message and invokes Job.Do.
func Do(ctx context.Context, message events.SQSMessage) (context.Context, error) {

	var m Message
	if err := json.Unmarshal([]byte(message.Body), &m); err != nil {
		return ctx, fmt.Errorf("error parsing message: %v", err)
	}
	if m.Job == "" {
		return ctx, errors.New("message invalid: 'job' required")
	}

	newJob, ok := jobs[m.Job]
	if !ok {
		return ctx, fmt.Errorf("job invalid: %q", m.Job)
	}

	return newJob().Do(ctx)
}

func init() {
	jobs = make(map[string]NewJob, 1)
	jobs[LogJobName] = NewLogJob
}
