package global

import "github.com/cloudwego/hertz/pkg/protocol/consts"

var (
	ParameterValidationError  = NewUException(consts.StatusBadRequest, "Parameter Validation Error", ParameterValidationErrorCode)
	RequestParameterBindError = NewUException(consts.StatusBadRequest, "RequestParameter Bind Error", RequestParameterBindErrorCode)
	MultipartFormError        = NewUException(consts.StatusBadRequest, "MultipartForm Occurrence Error", MultipartFormErrorCode)
	MultipartFileOpenError    = NewUException(consts.StatusBadRequest, "MultipartForm File Open Error", MultipartFileOpenErrorCode)
	MultipartFileCloseError   = NewUException(consts.StatusBadRequest, "MultipartForm File Close Error", MultipartFileCloseErrorCode)
	RPCClientCallError        = NewUException(consts.StatusServiceUnavailable, "RPC Client Call Error", RPCClientCallErrorCode)
)

var (
	InvalidTokenOrUnauthorized = NewUException(consts.StatusUnauthorized, "Invalid Token Or Unauthorized", InvalidTokenOrUnauthorizedCode)
	TokenNotFound              = NewUException(consts.StatusUnauthorized, "Token Not Found", TokenNotFoundCode)
)

const (
	ParameterValidationErrorCode = 50100 + iota
	RequestParameterBindErrorCode
	MultipartFormErrorCode
	MultipartFileOpenErrorCode
	MultipartFileCloseErrorCode
	RPCClientCallErrorCode
)
const (
	InvalidTokenOrUnauthorizedCode = 41000 + iota
	TokenNotFoundCode
)
