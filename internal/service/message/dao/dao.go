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
	"tuzi-tiktok/rds"
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
	PreMsgTime  int64
}

func GetMessageList(ctx context.Context, q QueryOption) ([]*message.Message, error) {
	msgList := make([]*model.Message, 0)
	var err error

	logger.Debugf("pre_msg_time is %v", q.PreMsgTime)
	// 获取历史消息
	if q.PreMsgTime == 0 {
		logger.Debug("================> entry history msg")
		msgList, err = getHistoryMessageList(q)
		if err != nil || msgList == nil {
			return nil, err
		}
	} else {
		logger.Debug("================> get now msg")
		msgList, err = getMessageListWithPreTime(q)
		if err != nil {
			logger.Errorf("get message listFrom failed , err : %v", err)
			return nil, err
		}
	}

	messageList := mMessage2kMessageMore(msgList)
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
	rKey := getHistoryMessageRedisKey(q.ToUid, q.Uid)
	exists := existsRedisKey(rKey)
	c := 0
	if exists == 1 {
		logger.Infof("%v key is exists", rKey)
		c, _ = rds.IRC.Get(context.Background(), rKey).Int()
	}
	rds.IRC.Set(ctx, rKey, c+1, time.Hour*24*7).Err()

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

func findMsgList(toUid, fromUid int64, t time.Time) ([]*model.Message, error) {
	return qMessage.Where(qMessage.ToUserID.Eq(toUid), qMessage.FormUserID.Eq(fromUid)).
		Where(qMessage.CreatedAt.Gt(t)).
		// Order(qMessage.CreatedAt).
		Find()
}

// getHistoryMessageList 获得历史消息列表
func getHistoryMessageList(q QueryOption) ([]*model.Message, error) {
	preKey := getPreMessageTimeRedisKey(q.Uid, q.ToUid)
	hKey := getHistoryMessageRedisKey(q.Uid, q.ToUid)
	if !HaveHistoryMessage(hKey) {
		return nil, nil
	}
	var t time.Time
	// 初次 || 再次 打开聊天，获取上次的聊天记录
	strTime, err := rds.IRC.Get(context.Background(), preKey).Result()
	if err != nil {
		logger.Errorf("get redis key %v error, err : %v", preKey, err)
		return nil, err
	}
	// t = 获取上次聊天的消息时间
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)
	logger.Debugf("get value from %v is %v ", preKey, t)
	// 获取双发的信息
	msgListFrom, _ := findMsgList(q.Uid, q.ToUid, t)
	msgListTo, _ := findMsgList(q.ToUid, q.Uid, t)

	msgListFrom = append(msgListFrom, msgListTo...)

	st := time.Now().Local().Format("2006-01-02 15:04:05")

	rds.IRC.Set(context.Background(), hKey, 0, time.Hour*24*7).Err()
	rds.IRC.Set(context.Background(), preKey, st, time.Hour*24*7).Err()
	return msgListFrom, nil
}

// getMessageListWithPreTime 获取消息列表
func getMessageListWithPreTime(q QueryOption) ([]*model.Message, error) {
	preKey := getPreMessageTimeRedisKey(q.Uid, q.ToUid)

	t := time.Unix(q.PreMsgTime, 0)

	mList, err := findMsgList(q.Uid, q.ToUid, t)
	if err != nil {
		logger.Errorf("get message listFrom failed , err : %v", err)
		return nil, err
	}

	newMsgTime := time.Now().Format("2006-01-02 15:04:05")

	err = rds.IRC.Set(context.Background(), preKey, newMsgTime, 0).Err()
	if err != nil {
		logger.Errorf("redis set %v error :%v", preKey, err)
		return nil, err
	}
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

// HaveHistoryMessage 从redis中读取，判断是否有新的消息
func HaveHistoryMessage(hKey string) bool {
	// hKey := getHistoryMessageRedisKey(toUid, fromUid)
	exists := existsRedisKey(hKey)
	if exists != 1 {
		logger.Infof("hkey = %v is no exists", hKey)
		return false
	}

	count, err := rds.IRC.Get(context.Background(), hKey).Int()
	if err != nil {
		logger.Errorf("get the value of %v error: %v", hKey, err)
		return false
	}
	logger.Debugf("HaveHistoryMessage hKey is %v, get count is %v", hKey, count)

	if count == 0 {
		logger.Infof("no history message from %v", hKey)
		return false
	}

	return true
}

func existsRedisKey(rKey string) int {
	exists, err := rds.IRC.Exists(context.Background(), rKey).Result()
	if err != nil {
		logger.Errorf("redis is wrong, err : %v", rKey, err)
		return 0
	}
	return int(exists)
}

// func getNewMessageCountRedisKey(toUid, fromUid int64) string {
// 	return fmt.Sprintf("%v-%v-new-message-count", toUid, fromUid)
// }

func getPreMessageTimeRedisKey(toUid, fromUid int64) string {
	return fmt.Sprintf("%v-%v-pre-message-time", toUid, fromUid)
}

func getHistoryMessageRedisKey(toUid, fromUid int64) string {
	return fmt.Sprintf("%v-%v-history-message-count", toUid, fromUid)
}
