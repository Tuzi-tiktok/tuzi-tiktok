package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/message"
	kmessage "tuzi-tiktok/kitex/kitex_gen/message"
)

func ToMessage(k *kmessage.Message) *message.Message {
	if k == nil {
		return nil
	}
	return &message.Message{
		Id:         k.Id,
		ToUserId:   k.ToUserId,
		FromUserId: k.FromUserId,
		Content:    k.Content,
		CreateTime: k.CreateTime,
	}
}

func ToMessageActionResponse(k *kmessage.MessageActionResponse) *message.MessageActionResponse {
	if k == nil {
		return nil
	}
	return &message.MessageActionResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
	}
}

func ToMessageChatResponse(k *kmessage.MessageChatResponse) *message.MessageChatResponse {
	if k == nil {
		return nil
	}
	messageList := make([]*message.Message, len(k.MessageList))
	for i := range k.MessageList {
		messageList[i] = ToMessage(k.MessageList[i])
	}
	return &message.MessageChatResponse{
		StatusCode:  k.StatusCode,
		StatusMsg:   k.StatusMsg,
		MessageList: messageList,
	}
}
