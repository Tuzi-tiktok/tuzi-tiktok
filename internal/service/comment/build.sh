#!/usr/bin/env bash
RUN_NAME="comment-api"

mkdir -p bin
go build -o bin/${RUN_NAME} tuzi-tiktok/service/comment
