package access

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
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

const DebugRequestKey = "DebugRequestKey"

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
		dump, err = os.OpenFile(fmt.Sprintf("dumps/%v.dump", time.Now().Format("01-02-15:04:05")), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

	// Process Request
	value, ok := c.Get(DebugRequestKey)
	req := "Bind Error or The parameter is incorrect"
	if ok {
		rq, err := sonic.Marshal(value)
		if err != nil {
			logger.Warnf("DebugDump Request Error %v", err)
		}
		var r bytes.Buffer
		err = json.Indent(&r, rq, "", "\t")
		if err != nil {
			logger.Warnf("DebugDump %v", err)
		}
		req = r.String()
	}

	// Process Response
	body := c.Response.Body()
	var resp bytes.Buffer
	err := json.Indent(&resp, body, "", "\t")
	if err != nil {
		logger.Warnf("DebugDump %v", err)
		return
	}
	rs := strings.Join([]string{
		"=====================>\n",
		"   ----",
		c.FullPath(),
		"\n----------- request \n",
		req,
		"\n----------- response \n",
		resp.String(),
		"\n<=====================\n",
	}, "")
	_, err = dump.WriteString(rs)
	if err != nil {
		logger.Warnf("DebugDump %v", err)
		return
	}

}

func DebugRecordRequest(c *app.RequestContext, v any) {
	if dump == nil {
		return
	}
	c.Set(DebugRequestKey, v)
}
