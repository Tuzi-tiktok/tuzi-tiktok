package global

import "github.com/cloudwego/hertz/pkg/protocol/consts"

var (
	ParameterValidationError  = NewUException(consts.StatusBadRequest, "", 0)
	RequestParameterBindError = NewUException(consts.StatusBadRequest, "", 0)
	MultipartFormError        = NewUException(consts.StatusBadRequest, "MultipartForm Occurrence Error", 0)
	MultipartFileOpenError    = NewUException(consts.StatusBadRequest, "MultipartForm Occurrence Error", 0)
	MultipartFileCloseError   = NewUException(consts.StatusBadRequest, "MultipartForm Occurrence Error", 0)
	RPCClientCallError        = NewUException(consts.StatusServiceUnavailable, "", 0)
)
