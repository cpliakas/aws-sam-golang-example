
# These environment variables must be set for deployment to work.
S3_BUCKET := $(S3_BUCKET)
STACK_NAME := $(STACK_NAME)

# Common values used throughout the Makefile, not intended to be configured.
TEMPLATE = template.yaml
PACKAGED_TEMPLATE = packaged.yaml

.PHONY: test
test: install
	go test -v ./lambdautils
	go test -v ./service/api
	go test -v ./service/error
	go test -v ./service/worker

.PHONY: clean
clean:
	rm -f api error worker $(PACKAGED_TEMPLATE)

.PHONY: install
install:
	go get ./...

.PHONY: update
update:
	go get -u ./...

api: ./service/api/main.go
	go build -o api ./service/api

error: ./service/error/main.go
	go build -o error ./service/error

worker: ./service/worker/main.go
	go build -o worker ./service/worker

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 $(MAKE) api
	GOOS=linux GOARCH=amd64 $(MAKE) error
	GOOS=linux GOARCH=amd64 $(MAKE) worker

.PHONY: build
build: clean lambda

.PHONY: run
run: build
	sam local start-api

.PHONY: package
package: build
	sam package --template-file $(TEMPLATE) --s3-bucket $(S3_BUCKET) --output-template-file $(PACKAGED_TEMPLATE)

.PHONY: deploy
deploy: package
	sam deploy --stack-name $(STACK_NAME) --template-file $(PACKAGED_TEMPLATE) --capabilities CAPABILITY_IAM

.PHONY: teardown
teardown:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)