# AWS SAM Golang Example

An example API written in Golang using the Amazon Serverless Application Model (AWS SAM).

## Overview

Go is arguably one of the easiest languages to write a RESTful API in. With the
addition of [Go support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/)
coupled with maturity of tooling around the (AWS Serverless Application Model)[https://github.com/awslabs/serverless-application-model],
deploying the API to serverless infrastrucutre is becoming much simpler, too.
Thanks to the [APEX Gateway](github.com/apex/gateway) library, you can even
write APIs in a familiar manner without changing how the code is structured.

The purpose of this project is to give a slightly more complex example than the
"hello world" ones provided by Amazon to show how [Go's standard net/http](https://golang.org/pkg/net/http/)
package can play nicely in a serverless world with AWS [API Gateway](https://aws.amazon.com/api-gateway/)
and [Lambda](https://aws.amazon.com/lambda/). It also shows how you can use Go
and Amazon's tooling to develop and test your API locally within in this model.

## Prerequisites

* [An AWS account](https://aws.amazon.com/)
* [Golang](https://golang.org/doc/install)
* [Docker](https://docs.docker.com/install/#cloud)
* [Node.js](https://nodejs.org/en/download/)
* [AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
* [AWS SAM Local](https://github.com/awslabs/aws-sam-local#windows-linux-macos-with-npm-recommended)

## Installation

*TODO*

## Usage

*TODO*
