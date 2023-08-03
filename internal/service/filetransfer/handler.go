package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/oss"
)

func Transfer(c context.Context, ctx *app.RequestContext) {
	oss.Ping()
}
