#!/usr/bin/env bash
RUN_NAME="message-api"

mkdir -p output/bin
go build -o output/bin/${RUN_NAME} tuzi-tiktok/service/message

