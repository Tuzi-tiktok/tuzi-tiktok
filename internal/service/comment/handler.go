package main

import (
	"context"
	"gorm.io/gorm/clause"
	"log"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	rpcAuth "tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	comment "tuzi-tiktok/kitex/kitex_gen/comment"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/utils"
	consts "tuzi-tiktok/utils/consts/comment"
)

var (
	qVideo     = query.Q.Video
	qComment   = query.Q.Comment
	qUser      = query.Q.User
	authClient authinfoservice.Client
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

func init() {
	var err error
	authClient, err = utils.NewAuth()
	if err != nil {
		log.Fatal(err)
	}
}

// Comment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) Comment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	logger.Infof("comment on video: %v, content: %s", req.VideoId, req.CommentText)

	// check token & get uid
	claims, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		resp = &comment.CommentResponse{
			StatusCode: consts.CommentInvalidToken,
			StatusMsg:  &consts.CommentInvalidTokenMsg,
		}
		return resp, nil
	}
	uid := claims.Payload.UID

	// check video exist
	v, e := qVideo.WithContext(ctx).Where(qVideo.ID.Eq(req.VideoId)).Select().Find()
	if e != nil {
		logger.Errorf("failed to query video by id: %d, err: %v", req.VideoId, err)
		return nil, e
	}
	if len(v) == 0 {
		logger.Infof("video: %d not exist", req.VideoId)
		resp = &comment.CommentResponse{
			StatusCode: consts.CommentVideoNotExist,
			StatusMsg:  &consts.CommentVideoNotExistMsg,
		}
		return resp, nil
	}
	vid := v[0].ID

	// check action type
	if req.ActionType == 1 {
		// publish comment
		c := &model.Comment{
			UID:     uid,
			Vid:     vid,
			Content: *req.CommentText,
		}
		e := qComment.WithContext(ctx).Clauses(clause.Returning{}).Create(c)
		if e != nil {
			logger.Errorf("failed to create comment, err: %v", e)
			return nil, e
		}
		logger.Infof("comment: %d created", c.ID)

		info, e := authClient.GetUserInfo(ctx, &rpcAuth.UserInfoRequest{
			UserId: uid,
			Token:  req.Token,
		})
		if e != nil {
			logger.Errorf("failed to get user info, err: %v", e)
			return nil, e
		}
		if info.StatusCode != 0 {
			logger.Errorf("failed to get user info, status code: %d, status msg: %s", info.StatusCode, *info.StatusMsg)
			resp = &comment.CommentResponse{
				StatusCode: info.StatusCode,
				StatusMsg:  info.StatusMsg,
			}
			return resp, nil
		}

		resp = &comment.CommentResponse{
			StatusCode: consts.CommentSucceed,
			StatusMsg:  &consts.CommentSucceedMsg,
			Comment: &comment.Comment{
				Id:         c.ID,
				User:       info.User,
				Content:    c.Content,
				CreateDate: c.CreatedAt.Format("01-02"),
			},
		}

		return resp, nil

	} else if req.ActionType == 2 {
		// delete comment
		info, e := qComment.WithContext(ctx).Where(qComment.ID.Eq(*req.CommentId)).Delete()
		if e != nil {
			logger.Errorf("failed to delete comment, err: %v", e)
			return nil, e
		}
		if info.RowsAffected == 0 {
			logger.Infof("comment: %d not exist", *req.CommentId)
			// TODO
		}
		logger.Infof("comment: %d deleted", *req.CommentId)

		resp = &comment.CommentResponse{
			StatusCode: consts.CommentSucceed,
			StatusMsg:  &consts.CommentSucceedMsg,
		}

		return resp, nil

	} else {
		logger.Infof("unknown action type: %d", req.ActionType)
		resp = &comment.CommentResponse{
			StatusCode: consts.CommentUnknownAction,
			StatusMsg:  &consts.CommentUnknownActionMsg,
		}
		return resp, nil
	}
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
