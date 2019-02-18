package lambdautils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// NewSQS returns *sqs.SQS and ensures the QUEUE_URL environment variable is
// set. If not, the application logs an error message and exits.
func NewSQS() SQS {
	return sqs.New(session.Must(session.NewSession()))
}

// SQS is an interface that is compatible with *sqs.SQS. This enables swapping
// out *sqs.SQS for a mock to facilitate unit testing.
type SQS interface {
	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
}

// SendMessage sends a message to SQS and panics on any errors.
func SendMessage(svc SQS, input *sqs.SendMessageInput) (output *sqs.SendMessageOutput) {
	var err error
	output, err = svc.SendMessage(input)
	if err != nil {
		panic(fmt.Errorf("error sending sqs message: %v", err))
	}
	return
}

// DeleteMessage deletes a message from SQS and panics on any errors.
func DeleteMessage(svc SQS, receiptHandle string) (output *sqs.DeleteMessageOutput) {
	var err error

	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(QueueURL()),
		ReceiptHandle: aws.String(receiptHandle),
	}

	if output, err = svc.DeleteMessage(input); err != nil {
		panic(fmt.Errorf("error deleting sqs message: %v", err))
	}

	return
}
