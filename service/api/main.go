package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/apex/gateway"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/cpliakas/aws-sam-golang-example/job"
	"github.com/cpliakas/aws-sam-golang-example/lambdautils"
)

// ContentType contains the Content-Type header sent on all responses.
const ContentType = "application/json; charset=utf8"

// MessageResponse models a simple message responses.
type MessageResponse struct {
	Message string `json:"message"`
}

// JobResponse models the respose returned by the /job endpoint.
type JobResponse struct {
	MessageResponse
	JobID string `json:"job_id"`
}

// WelcomeMessageResponse is the response returned by the / endpoint.
var WelcomeMessageResponse = MessageResponse{"Welcome to the example API!"}

// JobMessage is the message sent in JobResponse responses.
var JobMessage = "Job sent to queue."

// RootHandler is a http.HandlerFunc for the / endpoint.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(WelcomeMessageResponse)
}

// JobHandler implements http.Handler for the /job endpoint.
type JobHandler struct {
	svc sqsiface.SQSAPI
}

// JobHandlerFunc returns a http.HandlerFunc for the /job endpoint.
func JobHandlerFunc(svc sqsiface.SQSAPI) http.HandlerFunc {
	jh := JobHandler{svc: svc}
	return jh.ServeHTTP
}

// JobHandler implements http.Handler for the /job endpoint.
func (h JobHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lambdautils.Mustenv(lambdautils.EnvQueueURL)
	response := sendJobMessage(h.svc, job.LogJobName)
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)
}

// sendJobMessage sends an encoded job.Message to SQS.
func sendJobMessage(svc sqsiface.SQSAPI, name string) (response JobResponse) {

	output := lambdautils.SendMessage(svc, &sqs.SendMessageInput{
		QueueUrl:    aws.String(lambdautils.QueueURL()),
		MessageBody: aws.String(job.NewMessage(name).String()),
	})

	response = JobResponse{JobID: *output.MessageId}
	response.Message = JobMessage
	return
}

// RegisterRoutes registers the API's routes.
func RegisterRoutes() {
	sess := session.Must(session.NewSession())
	svc := sqs.New(sess)

	http.Handle("/", h(RootHandler))
	http.Handle("/job", h(JobHandlerFunc(svc)))
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("event: %v", r)
		w.Header().Set("Content-Type", ContentType)
		next.ServeHTTP(w, r)
	})
}

func main() {
	RegisterRoutes()
	log.Fatal(gateway.ListenAndServe(":3000", nil))
}
