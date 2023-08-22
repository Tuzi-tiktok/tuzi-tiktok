package dao

import (
	"context"
	"strconv"
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
	Uid        int64
	ToUid      int64
	ActionType int32
	Content    string
	PreMsgTime int64
}

func GetMessageList(ctx context.Context, q QueryOption) ([]*message.Message, error) {

	logger.Debugf("pre_msg_time is %v", q.PreMsgTime)
	tid, fid := q.ToUid, q.Uid

	lastTime := time.UnixMilli(q.PreMsgTime)

	ca := qMessage.Where(qMessage.ToUserID.Eq(tid)).Where(qMessage.FormUserID.Eq(fid))
	cb := qMessage.Where(qMessage.ToUserID.Eq(fid)).Where(qMessage.FormUserID.Eq(tid))

	g := qMessage.Where(ca).Or(cb)

	messages, err := qMessage.Debug().WithContext(ctx).
		Where(qMessage.CreatedAt.Gt(lastTime)).Where(g).Order(qMessage.CreatedAt).Find()

	if err != nil {
		logger.Errorf("GetMessageList Db Error %v", err)
		return nil, err
	}
	return covert(messages), nil
}

func covert(l []*model.Message) []*message.Message {
	messages := make([]*message.Message, len(l))
	for i := range messages {
		m := l[i]
		t := strconv.FormatInt(m.CreatedAt.UnixMilli(), 10)
		var (
			fid int64
			ct  string
		)
		// TODO Add the database is not empty constraint
		if m.FormUserID != nil {
			fid = *m.FormUserID
		}
		if m.Content != nil {
			ct = *m.Content
		}

		messages[i] = &message.Message{
			Id:         m.ID,
			ToUserId:   m.ToUserID,
			FromUserId: fid,
			Content:    ct,
			CreateTime: &t,
		}
	}
	return messages
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
