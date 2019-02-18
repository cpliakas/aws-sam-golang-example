package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cpliakas/aws-sam-golang-example/lambdautils"
)

// ErrorHandler handles errors in the dead letter queue.
func ErrorHandler(ctx context.Context, event events.SQSEvent, svc lambdautils.SQS) error {
	for _, message := range event.Records {

		// Handle the error here.
		log.Println("error handled:", message.MessageId)

		// Delete the message from the dead letter queue.
		lambdautils.DeleteMessage(svc, message.ReceiptHandle)
	}
	return nil
}

func handler(ctx context.Context, event events.SQSEvent) error {
	lambdautils.Mustenv(lambdautils.EnvQueueURL)
	svc := lambdautils.NewSQS()
	return ErrorHandler(ctx, event, svc)
}

func main() {
	lambda.Start(handler)
}
