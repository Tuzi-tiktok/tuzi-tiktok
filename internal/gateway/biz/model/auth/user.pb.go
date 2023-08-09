package auth

type UserLoginRequest struct {
	Username string `json:"username" form:"username" query:"username"` // 登录用户名
	Password string `json:"password" form:"password" query:"password"` // 登录密码
}

type UserLoginResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
	UserId     int64   `json:"user_id" form:"user_id" query:"user_id"`             // 用户id
	Token      string  `json:"token" form:"token" query:"token"`                   // 用户鉴权token
}

type UserRegisterRequest struct {
	Username string `json:"username" form:"username" query:"username"` // 注册用户名，最长32个字符
	Password string `json:"password" form:"password" query:"password"` // 密码，最长32个字符
}
type UserRegisterResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
	UserId     int64   `json:"user_id" form:"user_id" query:"user_id"`             // 用户id
	Token      string  `json:"token" form:"token" query:"token"`                   // 用户鉴权token
}

type UserInfoRequest struct {
	UserId int64  `json:"user_id" form:"user_id" query:"user_id"` // 用户id
	Token  string `json:"token" form:"token" query:"token"`       // 用户鉴权token
}

type UserInfoResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
	User       *User   `json:"user" form:"user" query:"user"`                      // 用户信息
}

type User struct {
	Id              int64   `json:"id" form:"id" query:"id"`                                           // 用户id
	Name            string  `json:"name" form:"name" query:"name"`                                     // 用户名称
	FollowCount     *int64  `json:"follow_count" form:"follow_count" query:"follow_count"`             // 关注总数
	FollowerCount   *int64  `json:"follower_count" form:"follower_count" query:"follower_count"`       // 粉丝总数
	IsFollow        bool    `json:"is_follow" form:"is_follow" query:"is_follow"`                      // true-已关注，false-未关注
	Avatar          *string `json:"avatar" form:"avatar" query:"avatar"`                               //用户头像
	BackgroundImage *string `json:"background_image" form:"background_image" query:"background_image"` //用户个人页顶部大图
	Signature       *string `json:"signature" form:"signature" query:"signature"`                      //个人简介
	TotalFavorited  *int64  `json:"total_favorited" form:"total_favorited" query:"total_favorited"`    //获赞数量
	WorkCount       *int64  `json:"work_count" form:"work_count" query:"work_count"`                   //作品数量
	FavoriteCount   *int64  `json:"favorite_count" form:"favorite_count" query:"favorite_count"`       //点赞数量
}
