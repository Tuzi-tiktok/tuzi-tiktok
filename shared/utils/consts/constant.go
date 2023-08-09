package consts

const (
	InternalServerError = 500
	ServiceUnavailable  = 503
)

const (
	UserExisted  = 409
	UserNotExist = 404
	WrongPwd     = 401
	InvalidToken = 403
)
const (
	PublishUploadSnapShotError = 5002 + iota
	PublishListError
	PublishTargetUserNotExist
)

var (
	SuccessMsg                    = "Success"
	PublishListErrorMsg           = "PublishListError"
	PublishTargetUserNotExistMsg  = "PublishTargetUserNotExist"
	PublishUploadSnapShotErrorMsg = "PublishUploadSnapShotError"
	InvalidTokenMsg               = "InvalidToken"
)

const (
	Success = 0
)
