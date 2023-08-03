package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/oss"
	// Required Import
	_ "tuzi-tiktok/oss/it"
)

func Transfer(c context.Context, ctx *app.RequestContext) {
	oss.Ping()
}
