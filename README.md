# AWS SAM Golang Example

[![Build Status](https://travis-ci.org/cpliakas/aws-sam-golang-example.svg?branch=master)](https://travis-ci.org/cpliakas/aws-sam-golang-example)
[![Go Report Card](https://goreportcard.com/badge/github.com/cpliakas/aws-sam-golang-example)](https://goreportcard.com/report/github.com/cpliakas/aws-sam-golang-example)

An example API and Worker written in Golang using the Amazon Serverless
Application Model (AWS SAM).

## Overview

Go is arguably one of the easiest languages in which to write a RESTful API.
With the addition of [Go support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/)
coupled with the maturity of tooling around the [AWS Serverless Application Model](https://github.com/awslabs/serverless-application-model),
deploying Golang-based APIs to serverless infrastructure is becoming much more
straightforward, too. Thanks to the [APEX Gateway](https://github.com/apex/gateway),
you can even write APIs in a familiar manner without changing how the code is
structured.

The purpose of this project is to give a slightly more complicated example than
the "hello world" ones provided by Amazon with a toolchain that supports both
local development and deployment to AWS as well as design patterns that
facilitate unit testing.

## Prerequisites

* [An AWS account](https://aws.amazon.com/)
* [Golang](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install)
* [Node.js](https://nodejs.org/en/download/)
* [AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
* [SAM CLI](https://aws.amazon.com/serverless/sam/)
* [jq](https://stedolan.github.io/jq/) (optional)

## Installation

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get github.com/cpliakas/aws-sam-golang-example/...
cd $GOPATH/src/github.com/cpliakas/aws-sam-golang-example
```

## Usage

### Run the API Locally

:warning: Make sure to install all the [Prerequisites](#prerequisites). On Mac
OSX and Windows, ensure that the Docker VM is running.

```sh
GOARCH=amd64 GOOS=linux go build -o api ./service/api
sam local start-api
```

or ...

```sh
make run
```

You can now consume the API using your tool of choice. [HTTPie](https://httpie.org/)
is pretty awesome.

```sh
http localhost:3000/
```

```
HTTP/1.1 200 OK
Content-Length: 28
Content-Type: application/json; charset=utf8
Date: Sat, 03 Feb 2018 20:12:07 GMT

{
    "message": "Hello, world!"
}
```

### Run the Worker Locally

*TODO*

### Deploy to AWS

First, set the following environment variables replacing `<MY-BUCKET-NAME>` and
`<MY-STACK-NAME>` as appropriate:

```sh
export S3_BUCKET="<MY-BUCKET-NAME>"
export STACK_NAME="<MY-STACK-NAME>"
```

Now build, package, and deploy the application:

```sh
GOOS=linux GOARCH=amd64 go build -o api ./service/api
GOOS=linux GOARCH=amd64 go build -o error ./service/error
GOOS=linux GOARCH=amd64 go build -o worker ./service/worker

sam package --template-file template.yaml --s3-bucket $S3_BUCKET --output-template-file packaged.yaml
sam deploy --stack-name $STACK_NAME --template-file packaged.yaml --capabilities CAPABILITY_IAM
```

or ...

```sh
make deploy
```

### Consume the Endpoint

The API endpoint is captured in the CloudFormation stack's `Endpoint` output
key. Either view the output value via the AWS Management Console, or run the
following command assuming the [jq](https://stedolan.github.io/jq/) tool is
installed:

```sh
aws cloudformation describe-stacks --stack-name $STACK_NAME | jq -r '.Stacks[0].Outputs[0].OutputValue'
```

Again, [HTTPie](https://httpie.org/) is a pretty awesome tool.

### View AWS Logs

Run the following command to get the CloudWatch logs for the API.

```sh
sam logs -n Api --stack-name $STACK_NAME
```

Replace `Api` with `Worker` or `Error` to get logs for the Lambda functions in
those resources as well.

:warning: The `sam` tool will throw a nasty stack trace if you try to view the
logs before the Lambda function has been invoked. Only run this command after
you have made requests to the corresponding handlers.
