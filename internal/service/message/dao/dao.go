package dao

import (
	"context"
	"time"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/message"
	"tuzi-tiktok/logger"
)

var (
	qMessage = query.Message
	qUser    = query.User
)

type QueryOption struct {
	Uid         int64
	ToUid       int64
	Action_type int32
	Content     string
}

func GetMessageList(ctx context.Context, q QueryOption) ([]*message.Message, error) {

	mMessageListFrom, err := getMessageListBy(q.Uid, q.ToUid)
	if err != nil {
		logger.Errorf("get message listFrom failed , err : %v", err)
		return nil, err
	}
	mMessageListTo, err := getMessageListBy(q.ToUid, q.Uid)
	if err != nil {
		logger.Errorf("get message listTo failed , err : %v", err)
		return nil, err
	}
	mergeMessageList := append(mMessageListFrom, mMessageListTo...)
	messageList := mMessage2kMessageMore(mergeMessageList)

	return messageList, nil
}

func ActionMessage(ctx context.Context, q QueryOption) (bool, error) {
	msg := model.Message{
		ToUserID:   q.ToUid,
		FormUserID: &q.Uid,
		Content:    &q.Content,
	}
	err := qMessage.WithContext(ctx).Create(&msg)
	if err != nil {
		logger.Errorf("user %d action message to %d failed, err: %v", q.Uid, q.ToUid, err)
		return false, err
	}
	return true, nil
}

func mMessage2kMessageMore(mMsgList []*model.Message) (msgList []*message.Message) {
	msgList = make([]*message.Message, 0)
	for _, msg := range mMsgList {
		newMsg := new(message.Message)
		newMsg.Id = msg.ID
		newMsg.ToUserId = msg.ToUserID
		newMsg.FromUserId = *msg.FormUserID
		newMsg.Content = *msg.Content
		strTime := msg.CreatedAt.Format("2006-01-02 15:04:05")
		newMsg.CreateTime = &strTime
		msgList = append(msgList, newMsg)
	}
	return
}

// getMessageListBy 获取消息列表，返回一个月（默认）之内的消息
func getMessageListBy(toUid, fromUid int64) ([]*model.Message, error) {
	return qMessage.Where(qMessage.ToUserID.Eq(toUid), qMessage.FormUserID.Eq(fromUid)).
		Where(qMessage.DeletedAt.IsNull()).
		Where(qMessage.CreatedAt.Gt(time.Now().AddDate(0, -1, 0))).
		Find()
}

func IsUserExist(uid int64) bool {
	c, err := qUser.Where(qUser.ID.Eq(uid)).Count()
	if err != nil {
		logger.Errorf("get User info error : %v", err)
		return false
	}
	return c == 1
}
