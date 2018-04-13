# AWS SAM Golang Example

[![Build Status](https://travis-ci.org/cpliakas/aws-sam-golang-example.svg?branch=master)](https://travis-ci.org/cpliakas/aws-sam-golang-example)
[![Go Report Card](https://goreportcard.com/badge/github.com/cpliakas/aws-sam-golang-example)](https://goreportcard.com/report/github.com/cpliakas/aws-sam-golang-example)

An example API written in Golang using the Amazon Serverless Application Model (AWS SAM).

## Overview

Go is arguably one of the easiest languages in which to write a RESTful API.
With the addition of [Go support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/)
coupled with the maturity of tooling around the [AWS Serverless Application Model](https://github.com/awslabs/serverless-application-model),
deploying the API to serverless infrastructure is becoming much more
straightforward, too. Thanks to the [APEX Gateway](https://github.com/apex/gateway),
you can even write APIs in a familiar manner without changing how the code is
structured.

The purpose of this project is to give a slightly more complicated example than
the "hello world" ones provided by Amazon to show how [Go's standard net/http](https://golang.org/pkg/net/http/)
package can play nicely in a serverless world with AWS [API Gateway](https://aws.amazon.com/api-gateway/)
and [Lambda](https://aws.amazon.com/lambda/). It also shows how you can use Go
and Amazon's tooling to develop and test your API locally within in this model.

## Prerequisites

* [An AWS account](https://aws.amazon.com/)
* [Golang](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install)
* [Node.js](https://nodejs.org/en/download/)
* [AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
* [AWS SAM Local](https://github.com/awslabs/aws-sam-local#windows-linux-macos-with-npm-recommended)

## Installation

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/cpliakas/aws-sam-golang-example
```

## Usage

### Run the API Locally

:warning: Make sure to install all the [Prerequisites](#prerequisites). On Mac
OSX and Windows, ensure that the Docker VM is running.

Build the binary, and run the API locally:

```sh
GOOS=linux go build -o main
sam local start-api
```

You can now consume the API using your tool of choice. [HTTPie](https://httpie.org/)
is pretty awesome.

```sh
http localhost:3000/hello
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