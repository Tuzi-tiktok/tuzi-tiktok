package dao

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/message"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/redis"
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
	// mMessageListTo, err := getMessageListBy(q.ToUid, q.Uid)
	// if err != nil {
	// 	logger.Errorf("get message listTo failed , err : %v", err)
	// 	return nil, err
	// }
	// mergeMessageList := append(mMessageListFrom, mMessageListTo...)
	messageList := mMessage2kMessageMore(mMessageListFrom)

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
	rKey := getNewMessageCountRedisKey(q.ToUid, q.Uid)
	exists := existsRedisKey(rKey)
	c := 0
	if exists == 1 {
		logger.Debugf("%v key is exists", rKey)
		c, _ = redis.IRC.Get(context.Background(), rKey).Int()
	}
	logger.Debugf("count is %v", c+1)
	redis.IRC.Set(ctx, rKey, c+1, 0).Err()

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
		strTime := msg.CreatedAt.Unix()
		str := strconv.Itoa(int(strTime))
		// strTime := msg.CreatedAt.Format("2006-01-02 15:04:05")
		newMsg.CreateTime = &str
		msgList = append(msgList, newMsg)
	}
	return
}

// getMessageListBy 获取消息列表，返回一星期（默认）之内的消息
func getMessageListBy(toUid, fromUid int64) ([]*model.Message, error) {
	cKey := getNewMessageCountRedisKey(toUid, fromUid)
	tKey := getPreMessageTimeRedisKey(toUid, fromUid)

	exists := existsRedisKey(tKey)
	// 查询时间值 默认为七天前
	t := time.Now().AddDate(0, 0, -7)
	if exists == 1 {
		v, err := redis.IRC.Get(context.Background(), tKey).Result()
		if err != nil {
			logger.Errorf("get the value of %v error: %v", tKey, err)
			return nil, err
		}
		loc, _ := time.LoadLocation("Asia/Shanghai")
		t, _ = time.ParseInLocation("2006-01-02 15:04:05", v, loc)
	}

	mList, err := qMessage.Where(qMessage.ToUserID.Eq(toUid), qMessage.FormUserID.Eq(fromUid)).
		Where(qMessage.CreatedAt.Gt(t)).
		Find()
	if err != nil {
		logger.Errorf("get message listFrom failed , err : %v", err)
		return nil, err
	}

	newMsgTime := time.Now().Format("2006-01-02 15:04:05")
	logger.Debugf("new redis key = %v : value = %v", tKey, newMsgTime)

	err = redis.IRC.Set(context.Background(), tKey, newMsgTime, 0).Err()
	if err != nil {
		logger.Errorf("redis set %v error :%v", tKey, err)
		return nil, err
	}
	// 将消息数重置
	redis.IRC.Set(context.Background(), cKey, 0, 0).Err()
	return mList, nil
}

func IsUserExist(uid int64) bool {
	c, err := qUser.Where(qUser.ID.Eq(uid)).Count()
	if err != nil {
		logger.Errorf("get User info error : %v", err)
		return false
	}
	return c == 1
}

// HaveNewMessage 从redis中读取，判断是否有新的消息
func HaveNewMessage(toUid, fromUid int64) bool {
	rKey := getNewMessageCountRedisKey(toUid, fromUid)

	exists := existsRedisKey(rKey)
	if exists != 1 {
		logger.Infof("rkey = %v is no exists", rKey)
		return false
	}

	count, err := redis.IRC.Get(context.Background(), rKey).Result()
	if err != nil {
		logger.Errorf("get the value of %v error: %v", rKey, err)
		return false
	}
	logger.Debugf("HaveNewMessage rKey is %v, get count is %v", rKey, count)

	c, _ := strconv.ParseInt(count, 10, 32)

	if c == 0 {
		logger.Infof("no new message from %v", fromUid)
		return false
	}

	return true
}

func existsRedisKey(rKey string) int {
	exists, err := redis.IRC.Exists(context.Background(), rKey).Result()
	if err != nil {
		logger.Errorf("redis is wrong, err : %v", rKey, err)
		return 0
	}
	return int(exists)
}

func getNewMessageCountRedisKey(toUid, fromUid int64) string {
	return fmt.Sprintf("%v-%v-new-message-count", toUid, fromUid)
}

func getPreMessageTimeRedisKey(toUid, fromUid int64) string {
	return fmt.Sprintf("%v-%v-pre-message-time", toUid, fromUid)
}
