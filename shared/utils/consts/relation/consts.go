package relation

const (
	RelationSucceed      = 0
	RelationActionFailed = 500
)

var (
	RelationSucceedMsg = "success"
)

const (
	RelationCommonErrorMSg    = 500
	RelationGetFavorListError = 5001
)

var (
	RelationCommonErrorMsg    = "Internal Server Error"
	RelationFollowFailedMsg   = "FollowFail"
	FavorGetFavorListErrorMsg = "Failed to get the likes list"
	RelationActionFailedMsg   = "Relation Action Failed"
)
