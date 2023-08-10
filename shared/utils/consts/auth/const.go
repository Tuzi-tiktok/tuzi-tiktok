package consts

const (
	AuthSucceed = 0
)

var (
	AuthSucceedMsg = "success"
)

const (
	AuthCommonError  = 500
	AuthInvalidToken = 403
	AuthUserExisted  = 409
	AuthUserNotExist = 404
	AuthWrongPwd     = 401
)

var (
	AuthCommonErrorMsg  = "Internal Server Error"
	AuthInvalidTokenMsg = "Invalid token"
	AuthUserExistedMsg  = "User name has been registered"
	AuthUserNotExistMsg = "User does not exist"
	AuthWrongPwdMsg     = "Wrong password"
)
