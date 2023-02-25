package model

import "github.com/neverTanking/TiktokByGo/db"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type User struct {
	db.User
	FavoriteCount  int64  `json:"favorite_count"`  // 喜欢数
	FollowCount    int64  `json:"follow_count"`    // 关注总数
	FollowerCount  int64  `json:"follower_count"`  // 粉丝总数
	TotalFavorited string `json:"total_favorited"` // 获赞数量
	WorkCount      int64  `json:"work_count"`      // 作品数
	IsFollow       bool   `json:"is_follow"`
}

// 单个视频
type Video struct {
	db.Video
	Author        User  `json:"author"` //author信息都在这里面
	FavoriteCount int64 `json:"favorite_count"`
	CommentCount  int64 `json:"comment_count"`
	IsFavorite    bool  `json:"is_favorite"`
}

type Comment struct {
	db.Comment
	User_ User `json:"user"`
}
