#!/usr/bin/env bash
set -e
GOOS=linux go build
docker build -t ejacobson96/testzipsvr .
docker push ejacobson96/testzipsvr
go clean