#!/usr/bin/env bash
RUN_NAME="relation-api"

mkdir -p bin
go build -o bin/${RUN_NAME} tuzi-tiktok/service/relation

