package main

import (
	"context"
	message "tuzi-tiktok/kitex/kitex_gen/message"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/message/dao"
	consts "tuzi-tiktok/utils/consts/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// GetMessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageList(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	resp = new(message.MessageChatResponse)

	uid := parseToken(req.Token)
	if uid == 0 {
		resp.StatusCode = consts.MESSAGE_API_Uid_FAILED
		resp.StatusMsg = &consts.MESSAGE_UID_GET_FAILED_MSG
		resp.MessageList = nil
		return
	}

	msgList, err := dao.GetMessageList(ctx, dao.QueryOption{
		Uid:        uid,
		ToUid:      req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	msgTime := req.PreMsgTime
	logger.Info(msgTime)

	if err != nil {
		logger.Errorf("Could not get the message list, err : %v", err)
		resp.StatusCode = consts.MESSAGE_API_GET_LIST_FAILED
		resp.StatusMsg = &consts.MESSAGE_GET_LIST_MSG
		return
	}

	resp.MessageList = msgList
	resp.StatusCode = consts.MESSAGE_API_SUCCESS
	resp.StatusMsg = &consts.MESSAGE_SUCCESS_MSG

	return
}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {

	// 如果content是空的，无需发送消息
	if req.Content == "" || req.ActionType != 1 {
		return &message.MessageActionResponse{
			StatusCode: consts.MESSAGE_API_CONTENT_NULL,
			StatusMsg:  &consts.MESSAGE_CONTENT_NULL_MES,
		}, nil
	}

	uid := parseToken(req.Token)
	if uid == 0 {
		return &message.MessageActionResponse{
			StatusCode: consts.MESSAGE_API_Uid_FAILED,
			StatusMsg:  &consts.MESSAGE_UID_GET_FAILED_MSG,
		}, nil
	}
	// ok := dao.IsUserExist(req.ToUserId)
	// if !ok {
	// 	resp.StatusCode = consts.MESSAGE_API_TOUID_NO_EXIST
	// 	resp.StatusMsg = &consts.MESSAGE_USER_NO_EXIST_MSG
	// 	return
	// }

	isAction, err := dao.ActionMessage(ctx, dao.QueryOption{
		Uid:        uid,
		ToUid:      req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})

	if err != nil || !isAction {
		return &message.MessageActionResponse{
			StatusCode: consts.MESSAGE_API_ACTION_FAILED,
			StatusMsg:  &consts.MESSAGE_ACTION_FAILED_MSG,
		}, nil
	}

	return &message.MessageActionResponse{
		StatusCode: consts.MESSAGE_API_SUCCESS,
		StatusMsg:  &consts.MESSAGE_SUCCESS_MSG,
	}, nil
}

// 获取用户uid
func parseToken(token string) int64 {
	claims, err := secret.ParseToken(token)
	if err != nil {
		logger.Errorf("failed to parse token, err: %v", err)
		return 0
	}
	// 从token获取当前用户uid
	uid := claims.Payload.UID
	logger.Infof("success to get uid : %d ", uid)
	return uid
}
