#!/bin/bash
RUN_NAME=gateway-api
mkdir -p bin
go build -o bin/${RUN_NAME} tuzi-tiktok/gateway
