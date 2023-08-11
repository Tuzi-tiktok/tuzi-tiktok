#!/usr/bin/env bash
RUN_NAME="publish-api"

mkdir -p output/bin
go build -o output/bin/${RUN_NAME} tuzi-tiktok/service/publish

