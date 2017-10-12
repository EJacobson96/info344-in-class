#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t ejacobson96/testserver .
docker push ejacobson96/testserver
go clean