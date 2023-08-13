package dao

import (
	"tuzi-tiktok/dao/query"
	consts "tuzi-tiktok/utils/consts/feed"
)

// isQueryUidInVideo 根据uid状态判断是否添加“where AuthorID = ？” 条件
func isQueryUidInVideo(uid int64, v query.IVideoDo) query.IVideoDo {
	if uid == consts.NOUSERSTATE {
		return v
	}
	return v.Where(qVideo.AuthorID.Eq(uid))
}

// isQueryUidInFavorite 根据uid状态判断是否添加“where UID = ？” 条件
func isQueryUidInFavorite(uid int64, v query.IFavoriteDo) query.IFavoriteDo {
	if uid == consts.NOUSERSTATE {
		return v
	}
	return v.Where(qFavorite.UID.Eq(uid))
}

// isQueryUidInRelation 根据uid状态判断是否添加“where ID = ？” 条件
func isQueryUidInRelation(uid int64, v query.IRelationDo) query.IRelationDo {
	if uid == consts.NOUSERSTATE {
		return v
	}
	return v.Where(qRelation.ID.Eq(uid))
}

//func isQueryUidInUser()  {
//
//}
