package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Transfer(c context.Context, ctx *app.RequestContext) {
	//form, err := ctx.MultipartForm()
	//if err != nil {
	//	e := ctx.Error(err)
	//	logger.Error("MultipartForm Occurrence Error", e)
	//}
	//token, title := form.Value["token"][0], form.Value["title"][0]
}
func ExtractUserId(token string) (int, error) {
	return 1, nil
}
