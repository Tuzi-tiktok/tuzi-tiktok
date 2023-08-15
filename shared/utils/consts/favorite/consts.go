package favorite

const (
	FavorUnKnownAction    = 402
	FavorTokenParseFailed = 401
	FavorSucceed          = 0
	FavorGetListFailed    = 5002
)

var (
	FavorSucceedMsg = "success"
)

const (
	FavorCommonErrorMSg    = 500
	FavorGetFavorListError = 5001
)

var (
	FavorUnKnownActionMsg    = "Favor UnKnown Action"
	FavorTokenParseFailedMsg = "Favor Token Parse Failed"
	FavorCommonErrorMsg      = "Internal Server Error"
	FavorGetListFailedMsg    = "Failed to get favor list"
)
