package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/comment"
	pcomment "tuzi-tiktok/kitex/kitex_gen/comment"
)

func ToComment(p *pcomment.Comment) *comment.Comment {
	if p == nil {
		return nil
	}
	return &comment.Comment{
		Id:         p.Id,
		User:       ToUser(p.User),
		Content:    p.Content,
		CreateDate: p.CreateDate,
	}
}

func ToCommentResponse(p *pcomment.CommentResponse) *comment.CommentResponse {
	if p == nil {
		return nil
	}
	return &comment.CommentResponse{
		StatusCode: p.StatusCode,
		StatusMsg:  p.StatusMsg,
		Comment:    ToComment(p.Comment),
	}
}

func ToCommentListResponse(p *pcomment.CommentListResponse) *comment.CommentListResponse {
	if p == nil {
		return nil
	}
	commentList := make([]*comment.Comment, len(p.CommentList))
	for i := range p.CommentList {
		commentList[i] = ToComment(p.CommentList[i])
	}
	return &comment.CommentListResponse{
		StatusCode:  p.StatusCode,
		StatusMsg:   p.StatusMsg,
		CommentList: commentList,
	}
}
