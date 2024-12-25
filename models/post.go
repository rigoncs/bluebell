package models

import "time"

// 内存对齐
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// APIPostDetail 帖子详情接口的结构体
type APIPostDetail struct {
	AuthorName       string             `json:"author_name"` // 作者姓名
	*Post                               // 帖子详情
	*CommunityDetail `json:"community"` // 社区详情
}
