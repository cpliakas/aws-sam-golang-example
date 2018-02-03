#!/usr/bin/env bash

GOOS=linux go build -o main
sam local start-api
