#!/bin/bash

processName1="Builder.bin"

go mod tidy
GOOS=linux go build -o ${processName1} main.go
