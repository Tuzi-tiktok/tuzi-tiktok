package relation

const (
	RelationFollowFailed     = 403
	RelationUnKnownAction    = 402
	RelationTokenParseFailed = 401
	RelationSucceed          = 0
	RelationActionFailed     = 500
)

var (
	RelationSucceedMsg = "success"
)

const (
	RelationCommonErrorMSg    = 500
	RelationGetFavorListError = 5001
)

var (
	RelationFollowFailedMsg     = "Relation Follow Failed"
	RelationUnKnownActionMsg    = "Relation UnKnown Action"
	RelationTokenParseFailedMsg = "Relation Token Parse Failed"
	RelationCommonErrorMsg      = "Internal Server Error"
	FavorGetFavorListErrorMsg   = "Failed to get the likes list"
)
