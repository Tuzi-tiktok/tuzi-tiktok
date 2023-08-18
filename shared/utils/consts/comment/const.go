package consts

const (
	CommentSucceed = 0
)

var (
	CommentSucceedMsg = "success"
)

const (
	CommentCommonError   = 500
	CommentUnknownAction = 400
	CommentInvalidToken  = 403
	CommentVideoNotExist = 4040
	CommentNotExist      = 4041
)

var (
	CommentCommonErrorMsg   = "Internal Server Error"
	CommentUnknownActionMsg = "Unknown action type"
	CommentInvalidTokenMsg  = "Invalid token"
	CommentVideoNotExistMsg = "Video does not exist"
	CommentNotExistMsg      = "Comment does not exist"
)
