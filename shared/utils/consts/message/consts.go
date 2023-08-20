package message

const (
	MESSAGE_API_SUCCESS = 0

	MESSAGE_API_GET_LIST_FAILED = 70001

	MESSAGE_API_ACTION_FAILED = 70002

	MESSAGE_API_Uid_FAILED = 70003

	MESSAGE_API_TOUID_NO_EXIST = 70004

	MESSAGE_API_CONTENT_NULL = 70005

	MESSAGE_API_NEW_MESSAGE_NULL = 0

	MESSAGE_API_FAIL = 70099
)

var (
	MESSAGE_SUCCESS_MSG = "success"

	MESSAGE_GET_LIST_MSG = "get_message_list_failed"

	MESSAGE_ACTION_FAILED_MSG = "action message failed"

	MESSAGE_UID_GET_FAILED_MSG = "get uid failed"

	MESSAGE_CONTENT_NULL_MES = "message content is null"

	MESSAGE_USER_NO_EXIST_MSG = "to uid no exist"

	MESSAGE_NEW_MESSAGE_NULL_MSG = "no new message"

	MESSAGE_FAIL_MSG = "failed"
)
