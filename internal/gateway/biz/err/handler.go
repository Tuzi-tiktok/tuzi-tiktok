package err

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/cloudwego/hertz/pkg/app"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/logger"
)

const (
	code = "status_code"
	msg  = "status_msg"
)
const maxTraceId = 1000

// ErrorHandlerMiddleware Handle Global Error
func ErrorHandlerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		tid := fastrand.Intn(maxTraceId)
		logger.Debugf("T: %d Access %v", tid, ctx.FullPath())
		ctx.Next(c)
		err := ctx.Errors.Last()
		if err != nil && err.Err != nil {
			e := err.Err
			var exception *global.UniformException
			if !errors.As(e, &exception) {
				return
			}
			var ret = make(map[string]any, 2)
			switch exception.InternalStatusCode {
			case global.InvalidTokenOrUnauthorizedCode, global.TokenNotFoundCode:
				logger.Warnf("T: %d Required Token is Not Found or InvalidToken Handler :%v; msg :%v;", tid, exception.HandlerName, exception.StatusMessage)
				ret[msg] = "Authentication failure Check Your Token"
			default:
				logger.Warnf("T: %d Service Face %v Catch Internal Error is %v", tid, exception.HandlerName, exception.StatusMessage)
				ret[msg] = "Service Unavailable Check Server Log"
			}
			ret[code] = exception.InternalStatusCode
			ctx.JSON(exception.HttpStatusCode, ret)
		}

		DebugDump(tid, ctx)
	}
}

var (
	dump  *os.File
	mutex sync.Mutex
)

func init() {
	_, ok := os.LookupEnv("TUZI_DEBUG")
	if ok {
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

func DebugDump(tid int, c *app.RequestContext) {
	if dump == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	body := c.Response.Body()
	var buf bytes.Buffer
	err := json.Indent(&buf, body, "", "\t")
	if err != nil {
		logger.Debugf("T: %v %v", tid, err)
		return
	}
	rs := strings.Join([]string{
		"=====================>\n",
		strconv.Itoa(tid),
		"----",
		c.FullPath(),
		"\n",
		buf.String(),
		"\n<=====================\n",
	}, "")
	_, err = dump.WriteString(rs)
	if err != nil {
		logger.Debugf("T: %v %v", tid, err)
		return
	}

}
