#!/usr/bin/env bash
RUN_NAME="transfer-http-api"

mkdir -p bin
go build -o bin/${RUN_NAME} tuzi-tiktok/service/filetransfer

