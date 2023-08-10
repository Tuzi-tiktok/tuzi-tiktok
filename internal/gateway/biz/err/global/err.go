package global

import (
	"tuzi-tiktok/logger"
)

type UniformException struct {
	HttpStatusCode     int
	StatusMessage      string
	InternalStatusCode int
	TraceError         error
	HandlerName        string
}

func (u *UniformException) Error() string {
	return u.StatusMessage
}

func (u *UniformException) WithError(e error) *UniformException {
	logger.Errorf("Handler Name is: %v With Err :%v", u.HandlerName, e)
	u.TraceError = e
	return u
}

func (u *UniformException) WithWarn(e error) *UniformException {
	logger.Warnf("Handler Name is: %v With Err %v", u.HandlerName, e)
	u.TraceError = e
	return u
}

func (u *UniformException) WithHandler(e string) *UniformException {
	u.HandlerName = e
	return u
}

func NewUException(httpCode int, msg string, code int) *UniformException {
	return &UniformException{
		HttpStatusCode:     httpCode,
		StatusMessage:      msg,
		InternalStatusCode: code,
	}
}
