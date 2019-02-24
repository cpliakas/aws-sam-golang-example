package lambdautils

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// SendMessage sends a message to SQS and panics on any errors.
func SendMessage(svc sqsiface.SQSAPI, input *sqs.SendMessageInput) (output *sqs.SendMessageOutput) {
	var err error
	output, err = svc.SendMessage(input)
	if err != nil {
		panic(fmt.Errorf("error sending sqs message: %v", err))
	}
	return
}

// DeleteMessage deletes a message from SQS and panics on any errors.
func DeleteMessage(svc sqsiface.SQSAPI, receiptHandle string) (output *sqs.DeleteMessageOutput) {
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
