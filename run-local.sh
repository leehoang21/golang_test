#!/usr/bin/env bash

go get github.com/githubnemo/CompileDaemon
pwd

go mod download
go mod tidy

swag fmt -g ./cmd/http/main.go
swag init -g ./cmd/http/main.go

export PORT=8080
export PATH=mongodb://docker:docker@0.0.0.0:27017
export DB_USER=docker
export DB_PASS=docker
export DB_NAME=djm-master

dt=$(date '+%d/%m/%Y %H:%M:%S')
echo "Run dev $dt"

CompileDaemon -log-prefix=false -build="go build -i -x ./cmd/http/main.go" -command="./main" -exclude-dir=".git"  -exclude-dir=".idea" -exclude-dir="vendor" -color
