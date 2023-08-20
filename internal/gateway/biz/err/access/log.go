package access

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
	"os"
	"strings"
	"sync"
	"time"
	"tuzi-tiktok/logger"
)

var (
	dump  *os.File
	mutex sync.Mutex
)

func init() {
	_, ok := os.LookupEnv("TUZI_DEBUG")
	if ok {
		logger.Debugf("--> The debug dump mode is enabled")
		var err error
		err = os.MkdirAll("dumps", 0666)
		if err != nil {
			logger.Error(err)
			panic(err)
		}
		dump, err = os.OpenFile(fmt.Sprintf("dumps/%v.dump", time.Now().Format("2006-01-02-15:04:05")), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		mutex = sync.Mutex{}
		if err != nil {
			logger.Error(err)
			panic(err)
		}
	}
}

func ALogMiddleware() []app.HandlerFunc {

	return []app.HandlerFunc{
		func(c context.Context, ctx *app.RequestContext) {
			ctx.Next(c)
			DebugDump(ctx)
		},
		accesslog.New(accesslog.WithAccessLogFunc(func(ctx context.Context, format string, v ...interface{}) {
			logger.Debugf(format, v...)
		})),
	}
}

func DebugDump(c *app.RequestContext) {
	if dump == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	body := c.Response.Body()
	var buf bytes.Buffer
	err := json.Indent(&buf, body, "", "\t")
	if err != nil {
		logger.Warnf("DebugDump %v", err)
		return
	}
	rs := strings.Join([]string{
		"=====================>\n",
		"----",
		c.FullPath(),
		"\n",
		buf.String(),
		"\n<=====================\n",
	}, "")
	_, err = dump.WriteString(rs)
	if err != nil {
		logger.Warnf("DebugDump %v", err)
		return
	}

}
