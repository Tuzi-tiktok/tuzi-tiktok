package relation

const (
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
	RelationUnKnownActionMsg    = "Relation UnKnown Action"
	RelationTokenParseFailedMsg = "Relation Token Parse Failed Msg"
	RelationCommonErrorMsg      = "Internal Server Error"
	RelationFollowFailedMsg     = "FollowFail"
	FavorGetFavorListErrorMsg   = "Failed to get the likes list"
	RelationActionFailedMsg     = "Relation Action Failed"
)
