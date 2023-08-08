package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/segmentio/ksuid"
	"mime/multipart"
	"net/http"
	"path"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/oss"
	"tuzi-tiktok/service/filetransfer/model"
)

func Transfer(c context.Context, ctx *app.RequestContext) {
	r := model.TransResult{}
	form, err := ctx.MultipartForm()
	if err != nil {
		r.Ok = false
		ctx.JSON(http.StatusOK, r)
		logger.Error("MultipartForm Occurrence Error", err)
		return
	}
	file := form.File["data"][0]
	randomKey, ext := ksuid.New().String(), path.Ext(file.Filename)
	objectName := randomKey + ext
	f, err := file.Open()
	if err != nil {
		r.Ok = false
		ctx.JSON(http.StatusOK, r)
		logger.Error("MultipartForm File Open Error", err)
		return
	}

	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			logger.Error("MultipartForm File Close Error", err)
		}
	}(f)

	err = oss.PutObject(objectName, f)
	if err != nil {
		r.Ok = false
		ctx.JSON(http.StatusOK, r)
		logger.Error("Oss Put  File  Error", err)
		return
	}
	r.Ok, r.Url = true, oss.GetAddress(objectName)
	ctx.JSON(200, r)
}
