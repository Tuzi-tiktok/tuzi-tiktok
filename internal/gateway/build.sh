#!/bin/bash
RUN_NAME=gateway-api
mkdir -p output/bin
go build -o output/bin/${RUN_NAME} tuzi-tiktok/gateway
